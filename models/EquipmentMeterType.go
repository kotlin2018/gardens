package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type (
	EquipmentMeterType struct {
		Id          int       `form:"id"`
		MeterTypeNO string    `orm:"column(meter_type_no)"`
		MeterType   string    `orm:"column(meter_type)"`
		VendorNO    string    `orm:"column(vendor_no)"`
		PtAddress   string    `orm:"column(pt_address)"`
		CtAddress   string    `orm:"column(ct_address)"`
		ThreePhase  int       `orm:"column(three_phase)"`
		Used        int       `orm:"column(tag)"`
		CreateUser  string    `orm:"column(createuser)"`
		CreateDate  time.Time `orm:"column(createdate)"`
		ChangeUser  string    `orm:"column(changeuser)"`
		ChangeDate  time.Time `orm:"column(changedate)"`
	}

	EquipmentMeterTypeGrid struct {
		Id          int
		MeterTypeNO string
		MeterType   string
		VendorNO    string
		VendorDesc  string
		PtAddress   string
		CtAddress   string
		ThreePhase  int

		Used       int
		CreateUser string
		CreateDate time.Time
		ChangeUser string
		ChangeDate time.Time
	}

	EquipmentMeterTypeQueryParam struct {
		BaseQueryParam
		MeterTypeNO string
		MeterType   string
		Used        string //为空不查询，有值精确查询
	}
)

func init() {
	orm.RegisterModel(new(EquipmentMeterType))
}

func EquipmentMeterTypeTBName() string {
	return "equipment_meter_type"
}

func (this *EquipmentMeterType) TableName() string {
	return EquipmentMeterTypeTBName()
}

func EquipmentMeterTypeSelect(params *EquipmentMeterTypeQueryParam)([]*EquipmentMeterType, error){
	query := orm.NewOrm().QueryTable(EquipmentMeterTypeTBName())

	params.Limit = -1

	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}

	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	data := make([]*EquipmentMeterType, 0)
	_, err := query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	if(err != nil){
		return nil, err
	}
	return data, nil
}

func EquipmentMeterTypePageList(params *EquipmentMeterTypeQueryParam) ([]*EquipmentMeterTypeGrid, int64) {
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}

	if params.Order == "desc" {
		sortorder += " DESC"
	}

	query := orm.NewOrm()
	lists := make([]*EquipmentMeterTypeGrid, 0)

	sql := fmt.Sprintf(`SELECT mt.id, 
									    mt.meter_type_no meter_type_n_o, 
									    mt.meter_type, 
									    mt.vendor_no vendor_n_o, 
									    ev.vendor_desc, 
									    mt.pt_address, 
									    mt.ct_address, 
									    mt.three_phase, 
									    mt.tag used, 
									    mt.createuser create_user, 
									    mt.createdate create_date, 
									    mt.changeuser change_user, 
									    mt.changedate change_date
								   FROM equipment_meter_type AS mt 
								   LEFT JOIN equipment_vendor AS ev ON mt.vendor_no = ev.vendor_no 
                                  WHERE mt.meter_type_no LIKE '%s%%'
                                    AND mt.meter_type LIKE '%%%s%%'
								  ORDER BY %s								 
							     `,
		params.MeterTypeNO,
		params.MeterType,
		sortorder,
	)

	total, err := query.Raw(sql).QueryRows(&lists)
	if err != nil {
		return nil, 0
	}

	sql = sql + fmt.Sprintf(" LIMIT %d, %d", params.Offset, params.Limit)

	_, err = query.Raw(sql).QueryRows(&lists)
	if err != nil {
		return nil, 0
	}
	return lists, total
}

func EquipmentMeterTypeDataList(params *EquipmentMeterTypeQueryParam) [] *EquipmentMeterTypeGrid {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := EquipmentMeterTypePageList(params)
	return data
}

func EquipmentMeterTypeBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm()
	sql := "DELETE from equipment_meter_type WHERE id in(?)"
	res, err := query.Raw(sql, ids).Exec()
	if err != nil {
		return 0, err
	}

	num, _ := res.RowsAffected()
	return num, nil
}

func EquipmentMeterTypeOne(id int) (*EquipmentMeterType, error) {
	o := orm.NewOrm()
	m := EquipmentMeterType{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func EquipmentMeterTypeAdd(meterType *EquipmentMeterType) (int64, error) {
	id, err := orm.NewOrm().Insert(meterType)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *EquipmentMeterType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}
