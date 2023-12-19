package models

type UserModel struct {
	ID          string `gorm:"primaryKey;column:id_user" json:"id_user"`
	NamaLengkap string `json:"nama_lengkap"`
	UserName    string `gorm:"column:username" json:"username"`
	Password    string `json:"password"`
}
