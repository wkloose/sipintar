package models

import (
    "github.com/google/uuid"
)

type MaterialContent struct {
    ID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    MaterialID uuid.UUID  `gorm:"type:uuid"` // relasi ke Material
    Type       string     // "text", "video", "image"
    Text       *string    // hanya diisi jika Type = "text"
    Link       *string    // untuk video/image berbasis URL
    ImageBlob  *[]byte    // untuk menyimpan blob gambar jika perlu
    Order      int        // urutan konten
}
