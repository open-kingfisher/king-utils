package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/open-kingfisher/king-utils/common/log"
	"github.com/open-kingfisher/king-utils/config"
	"reflect"
	"strings"
	"time"
)

var DB *sql.DB

func init() {
	var err error
	for i := 1; i <= 64; i <<= 1 {
		DB, err = sql.Open("mysql", config.DBURL)
		if err != nil {
			log.Errorf("Connect %s failed: %s; on...: %v", config.DBURL, err, i)
			time.Sleep(time.Second * time.Duration(i))
			continue
		} else {
			break
		}
	}
	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(0)
	for i := 1; i <= 64; i <<= 1 {
		if err := DB.Ping(); err != nil {
			log.Errorf("Connect %s fatal %s", config.DBURL, err)
			time.Sleep(time.Second * time.Duration(i))
			continue
		} else if i == 64 {
			log.Fatal(err)
		} else {
			break
		}
	}
}

func checkTable(table string) error {
	// CREATE TABLE IF NOT EXISTS test (id INT NOT NULL AUTO_INCREMENT, data JSON NOT NULL, PRIMARY KEY (id));
	sql := "CREATE TABLE IF NOT EXISTS " + table + " (id INT NOT NULL AUTO_INCREMENT, data JSON NOT NULL, PRIMARY KEY (id))"
	if _, err := DB.Exec(sql); err != nil {
		return err
	}
	return nil
}

func ToJSON(obj interface{}) (string, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	replacer := strings.NewReplacer("\\n", "\\\\n", "\\r", "\\\\r", "\"", "\\\"", "'", "\\'", "\\", "\\\\")

	return replacer.Replace(string(jsonStr)), nil
}

func Insert(table string, obj interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if err := InsertTx(tx, table, obj); err != nil {
		tx.Rollback()
		log.Errorf("db insert error:%s; table:%s; object:%+v", err, table, obj)
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func InsertTx(tx *sql.Tx, table string, obj interface{}) error {
	//if err := checkTable(table); err != nil {
	//	return err
	//}
	strValue, err := ToJSON(obj)
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO " + table + " (data) VALUES ('" + strValue + "')")
	return err
}

func InsertIfNotExist(table, id string, obj interface{}) error {
	var o interface{}
	if GetById(table, id, &o) == sql.ErrNoRows {
		return Insert(table, obj)
	}
	return nil
}

// Update||Insert
func Upsert(table, id string, obj interface{}) error {
	var o interface{}
	err := GetById(table, id, &o)
	if err == sql.ErrNoRows {
		return Insert(table, obj)
	}
	return Update(table, id, obj)
}

func GetById(table string, id string, obj interface{}) error {
	return Get(table, map[string]interface{}{"$.id": id}, obj)
}

// https://dev.mysql.com/doc/refman/5.7/en/json.html#json-paths
func Get(table string, kvs map[string]interface{}, obj interface{}) error {
	//if err := checkTable(table); err != nil {
	//	return err
	//}
	union := ""
	args := make([]interface{}, 0)
	for k, value := range kvs {
		union = union + "AND data->'" + k + "'=? "
		switch v := value.(type) {
		case string:
			args = append(args, v)
		case int, int32, int64, uint, uint32, uint64:
			args = append(args, fmt.Sprint(v))
		default:
			return fmt.Errorf("sql query value unknown type: %+v", v)
		}
	}
	query := "SELECT data FROM " + table + " WHERE " + strings.TrimPrefix(union, "AND")
	jsonStr := ""
	if err := DB.QueryRow(query, args...).Scan(&jsonStr); err == nil {
		if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func List(field, table string, result interface{}, clause string, args ...interface{}) error {
	//if err := checkTable(table); err != nil {
	//	return err
	//}
	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}
	sliceValue := resultValue.Elem()
	elem := sliceValue.Type().Elem()
	query := fmt.Sprintf("SELECT %s FROM %s", field, table)

	if clause != "" {
		query = query + " " + clause
	}
	log.Infof("%s %v", query, args)
	// 避免SQL注入
	// db.Query("SELECT name FROM users WHERE age=?", age)
	rows, err := DB.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		jsonStr := ""
		err := rows.Scan(&jsonStr)
		if err != nil {
			return err
		}
		elemp := reflect.New(elem)
		json.Unmarshal([]byte(jsonStr), elemp.Interface())
		sliceValue = reflect.Append(sliceValue, elemp.Elem())
		i++
	}
	resultValue.Elem().Set(sliceValue.Slice(0, i))
	return nil
}

func Update(table, id string, newObj interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	if err := UpdateTx(tx, table, id, newObj); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func UpdateTx(tx *sql.Tx, table, id string, newObj interface{}) error {
	if err := DeleteTx(tx, table, id); err != nil {
		return err
	}
	if err := InsertTx(tx, table, newObj); err != nil {
		return err
	}
	return nil
}

//string|int value works for now
func UpdateKVS(table, id string, kvs map[string]interface{}) error {
	union := ""
	for key, v := range kvs { //key should be [json-path], e.g:$.id
		switch value := v.(type) {
		case string:
			union = union + ",'" + key + "','" + value + "'"
		case int, int32, int64, uint, uint32, uint64:
			union = union + ",'" + key + "'," + fmt.Sprint(value)
		default:
			return fmt.Errorf("unknown type: %+v", v)
		}
	}
	sql := "UPDATE " + table + " SET data=" + "JSON_SET(data" + union + ") WHERE data->'$.id'='" + id + "'"
	_, err := DB.Exec(sql)
	return err
}

func Delete(table, id string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if err := DeleteTx(tx, table, id); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteTx(tx *sql.Tx, table, id string) error {
	sql := "DELETE FROM " + table + " WHERE data->'$.id'='" + id + "'"
	_, err := tx.Exec(sql)
	return err
}

func UpdateUUId(table, uuId string, kvs map[string]interface{}) error {
	union := ""
	for key, v := range kvs { //key should be [json-path], e.g:$.id
		switch value := v.(type) {
		case string:
			union = union + ",'" + key + "','" + value + "'"
		case int, int32, int64, uint, uint32, uint64:
			union = union + ",'" + key + "'," + fmt.Sprint(value)
		default:
			log.Warnf("unknown type: %+v", v)
		}
	}
	sql := "UPDATE " + table + " SET data=" + "JSON_SET(data" + union + ") WHERE data->'$.uu_id'='" + uuId + "'"
	_, err := DB.Exec(sql)
	return err
}
