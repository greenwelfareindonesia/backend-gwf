package ecopedia

import "greenwelfare/user"

type EcopediaInput struct {
	Judul     string `form:"Judul" binding:"required"`
	Subjudul  string `form:"SubJudul"`
	Deskripsi string `form:"Deskripsi" binding:"required"`
	Srcgambar string `form:"SrcGambar"`
	Referensi string `form:"Referensi"`
}

type UserActionToEcopedia struct {
	Comment string `form:"Comment"`
	IsLike  bool   `form:"Is_like"`
	User    user.User 
}

type EcopediaID struct {
	ID int `uri:"ID" binding:"required"`
}
