package model

type Bill struct {
    Base
    Io                          string                      `json:"io" binding:"required"`
    Uid                         uint                        `json:"uid"`
    Tid                         uint                        `json:"tid" binding:"required"`
    Tag                         string                      `gorm:"size:200" json:"tag" binding:"required"`
    Money                       string                      `json:"money" binding:"required"`
    Remark                      string                      `gorm:"size:500" json:"remark" binding:"required"`
    DateTime                    string                      `gorm:"type:datetime" json:"datetime" binding:"required"`
}

func (*Bill) TableName() string {

    return "asset_bill"
}

func (*Bill) Create(mbill Bill) (uint) {
    DbWrite.NewRecord(mbill)

    DbWrite.Create(&mbill)

    DbWrite.NewRecord(mbill)

    return mbill.Id
}

func (*Bill) FindAll() ([]Bill) {
    var Mbills []Bill
    var Where = &Bill{}

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").Find(&Mbills)

    return Mbills
}

func (*Bill) FindSetByUid(uid uint) ([]Bill) {
    var Mbills []Bill
    var Where = &Bill{ Uid: uid }

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").Find(&Mbills)

    return Mbills
}

func (*Bill) FindOneByIdUid(mbill Bill) (Bill) {
    var Mbill Bill
    var Mbase = Base{ Id: mbill.Id }
    var Where = &Bill{ Base: Mbase, Uid: mbill.Uid }

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").First(&Mbill)

    return Mbill
}

func (*Bill) SaveOneByIdUid(mbill Bill) (int64) {
    var Mbase = Base{ Id: mbill.Id }
    var Where = &Bill{ Base: Mbase, Uid: mbill.Uid }

    return DbWrite.Model(Bill{}).Where(Where).Where("`status` = 0").Updates(mbill).RowsAffected
}
