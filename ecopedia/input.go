package ecopedia

import "greenwelfare/user"

type EcopediaInput struct {
	Judul     string `form:"judul" binding:"required"`
	Subjudul  string `form:"subjudul"`
	Deskripsi string `form:"Deskripsi" binding:"required"`
	Srcgambar string `form:"src_gambar"`
	Referensi string `form:"referensi"`
}

type UserActionToEcopedia struct {
	Comment string `form:"comment"`
	IsLike  bool   `form:"is_like"`
	User    user.User 
}

type EcopediaID struct {
	ID int `uri:"id" binding:"required"`
}
