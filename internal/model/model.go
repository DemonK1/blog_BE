package model

// Model 创建公共Model
type Model struct {
  ID         uint32 `gorm:"primary_key" json:"id"`
  CreatedOn  uint32 `json:"created_on"`
  ModifiedOn uint32 `json:"modified_on"`
  DeleteOn   uint32 `json:"delete_on"`
  IsDel      uint8  `json:"is_del"`
  CreateBy   string `json:"create_by"`
  ModifiedBy string `json:"modified_by"`
}
