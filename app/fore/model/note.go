package model

type Note struct {
    Base
    Uid                         uint                        `json:"uid"`
    Tid                         uint                        `json:"tid" binding:"required"`
    Tag                         string                      `gorm:"size:200" json:"tag" binding:"required"`
    Title                       string                      `gorm:"size:500" json:"title" binding:"required"`
    Content                     string                      `gorm:"size:50000000" json:"content" binding:"required"`
    DateTime                    string                      `gorm:"type:datetime" json:"datetime" binding:"required"`
}

func (*Note) TableName() string {

    return "asset_note"
}

func (*Note) Create(mnote Note) (uint) {
    DbWrite.NewRecord(mnote)

    DbWrite.Create(&mnote)

    DbWrite.NewRecord(mnote)

    return mnote.Id
}

func (*Note) FindAll() ([]Note) {
    var Mnotes []Note
    var Where = &Note{}

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").Find(&Mnotes)

    return Mnotes
}

func (*Note) FindSetByUid(uid uint) ([]Note) {
    var Mnotes []Note
    var Where = &Note{ Uid: uid }

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").Find(&Mnotes)

    return Mnotes
}

func (*Note) FindOneByIdUid(mnote Note) (Note) {
    var Mnote Note
    var Mbase = Base{ Id: mnote.Id }
    var Where = &Note{ Base: Mbase, Uid: mnote.Uid }

    DbQuery.Where(Where).Where("`status` = 0").Order("date_time desc").First(&Mnote)

    return Mnote
}

func (*Note) SaveOneByIdUid(mnote Note) (int64) {
    var Mbase = Base{ Id: mnote.Id }
    var Where = &Note{ Base: Mbase, Uid: mnote.Uid }

    return DbWrite.Model(Note{}).Where(Where).Where("`status` = 0").Updates(mnote).RowsAffected
}
