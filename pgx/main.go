package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM data_types ")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var smallserial_field pgtype.Int4
		var serial_field int
		var bigserial_field int
		var smallint_field int
		var integer_field int
		var bigint_field int
		var numeric_field int
		var real_field float32
		var double_field float32
		var varchar_field string
		var char_field string
		var text_field string
		var bytea_field []byte
		var timestamp_field time.Time
		var timestamptz_field time.Time
		var date_field time.Time
		var time_field pgtype.Time
		var timetz_field pgtype.Time
		var interval_field pgtype.Interval
		var bool_filed bool
		var point_filed pgtype.Point
		var line_field pgtype.Line
		var lseg_field pgtype.Lseg
		var box_field pgtype.Box
		var path_field pgtype.Path
		var polygon_field pgtype.Polygon
		var circle_field pgtype.Circle
		var cidr_field pgtype.Inet //27
		var inet_field pgtype.Inet
		var macaddr_field string
		var macaddr8_field string
		var bit_field pgtype.Bit
		var tsvector_field string
		var tsquery_field string
		var uuid_field pgtype.UUID
		var xml_type string
		var json_field pgtype.JSON
		var jsonb_field pgtype.JSONB
		var jsonpath_field string
		var intrange_field pgtype.Int4range
		var numrange_field pgtype.Numrange
		var int8range_field pgtype.Int8range
		var tsrange_field pgtype.Tsrange
		var tstzrange_field pgtype.Tstzrange
		var datarange_field pgtype.Daterange
		var intarray_field pgtype.Int4Array
		var strarray_field pgtype.VarcharArray
		err := rows.Scan(&smallserial_field, &serial_field, &bigserial_field, &smallint_field, &integer_field, &bigint_field, &numeric_field, &real_field, &double_field, &varchar_field, &char_field, &text_field, &bytea_field, &timestamp_field, &timestamptz_field, &date_field, &time_field, &timetz_field, &interval_field, &bool_filed, &point_filed, &line_field, &lseg_field, &box_field, &path_field, &polygon_field, &circle_field, &cidr_field, &inet_field, &macaddr_field, &macaddr8_field, &bit_field, &tsvector_field, &tsquery_field, &uuid_field, &xml_type, &json_field, &jsonb_field, &jsonpath_field, &intrange_field, &numrange_field, &int8range_field, &tsrange_field, &tstzrange_field, &datarange_field, &intarray_field, &strarray_field)
		f := fmt.Sprintf("error:%v,smallserial_field:%v,serial_field:%v,bigserial_field:%v,smallint_field:%v,integer_field:%v,bigint_field:%v,numeric_field:%v,real_field:%v,double_field:%v,varchar_field:%v,char_field:%v,text_field:%v,bytea_field:%v,timestamp_field:%v,timestamptz_field:%v,date_field:%v,time_field:%v,timetz_field:%v,interval_field:%v,bool_filed:%v,point_filed:%v,line_field:%v,lseg_field:%v,box_field:%v,path_field:%v,polygon_field:%v,circle_field:%v,cidr_field:%v,inet_field:%v,macaddr_field:%v,macaddr8_field:%v,bit_field:%v,tsvector_field:%v,tsquery_field:%v,uuid_field:%v,xml_type:%v,json_field:%v,jsonb_field:%v,jsonpath_field:%v,intrange_field:%v,numrange_field:%v,int8range_field:%v,tsrange_field:%v,tstzrange_field:%v,datarange_field:%v,intarray_field:%v,strarray_field:%v", err, smallserial_field, serial_field, bigserial_field, smallint_field, integer_field, bigint_field, numeric_field, real_field, double_field, varchar_field, char_field, text_field, bytea_field, timestamp_field, timestamptz_field, date_field, time_field, timetz_field, interval_field, bool_filed, point_filed, line_field, lseg_field, box_field, path_field, polygon_field, circle_field, cidr_field, inet_field, macaddr_field, macaddr8_field, bit_field, tsvector_field, tsquery_field, uuid_field, xml_type, json_field, jsonb_field, jsonpath_field, intrange_field, numrange_field, int8range_field, tsrange_field, tstzrange_field, datarange_field, intarray_field, strarray_field)
		fmt.Println(f)
		var ll []int
		intarray_field.AssignTo(&ll)
		fmt.Println(ll)
	}
}
