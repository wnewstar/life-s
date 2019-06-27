package logic

import (
    "time"
    "errors"
    "strconv"
    "encoding/base64"
    "golang.org/x/crypto/bcrypt"
    ApiModel "life/app/api/model"
)

type User struct {}

func (*User) Login(user ApiModel.User) (back map[string]map[string]interface{}, err error) {
    var Time uint64
    var Muser ApiModel.User

    err = errors.New("无效用户")
    Time = uint64(time.Now().Unix())
    Muser = Muser.FindOneByUsername(user.Username)

    if (Muser.Id > 0) {
        if err = bcrypt.CompareHashAndPassword([]byte(Muser.Password), []byte(user.Password)); err == nil {
            ApiModel.Musers[Muser.Id] = Muser
            a := uint64(Muser.Id)
            b := Time + 7200
            back = make(map[string]map[string]interface{})
            back["user"] = make(map[string]interface{})
            back["auth"] = make(map[string]interface{})
            back["user"]["id"] = Muser.Id
            back["user"]["uname"] = Muser.Nickname
            back["user"]["ctime"] = Muser.TimeCreate
            back["auth"]["etime"] = Time + 7200
            back["auth"]["token"] = base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(a, 10) + "|" + strconv.FormatUint(b, 10)))
        }
    }

    return back, err
}
