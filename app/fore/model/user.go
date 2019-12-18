package model

type User struct {
    Base
    Nickname                    string                      `gorm:"size:50" json:"nickname"`
    Username                    string                      `gorm:"size:50" json:"username" binding:"required"`
    Password                    string                      `gorm:"size:200" json:"password" binding:"required"`
}

func (*User) TableName() (string) {

    return "asset_user"
}

func (*User) Create(muser User) (uint) {
    DbWrite.NewRecord(muser)

    DbWrite.Create(&muser)

    DbWrite.NewRecord(muser)

    return muser.Id
}

func (*User) FindAll() ([]User) {
    var Musers []User

    DbQuery.Where("status = 0").Find(&Musers)

    return Musers
}

func (*User) FindOneById(id uint) (User) {
    var Muser User
    var Where = &Base{ Id: id }

    DbQuery.Where(Where).Where("status = 0").First(&Muser)

    return Muser
}

func (*User) FindOneByUsername(name string) (User) {
    var Muser User
    var Where = &User{ Username: name }

    DbQuery.Where(Where).Where("status = 0").First(&Muser)

    return Muser
}

func (*User) SaveOneById(muser User) (int64) {
    var Where = &Base{ Id: muser.Id }

    return DbWrite.Where(Where).Where("status = 0").Updates(muser).RowsAffected
}

func (*User) SaveOneByUsername(muser User) (int64) {
    var Where = &User{ Username: muser.Username }

    return DbWrite.Model(User{}).Where(Where).Where("status = 0").Updates(muser).RowsAffected
}
