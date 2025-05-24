package service

import (
	"context"
	"sync"

	conf "github.com/CocaineCong/gin-mall/config"
	"github.com/CocaineCong/gin-mall/consts"
	util "github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/types"
)

var CategorySrvIns *CategorySrv
var CategorySrvOnce sync.Once

type CategorySrv struct {
}

func GetCategorySrv() *CategorySrv {
	CategorySrvOnce.Do(func() {
		CategorySrvIns = &CategorySrv{}
	})
	return CategorySrvIns
}

// CategoryList 列举分类
func (s *CategorySrv) CategoryList(ctx context.Context, req *types.ListCategoryReq) (resp interface{}, err error) {
	categories, err := dao.NewCategoryDao(ctx).ListCategory()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	cResp := make([]*types.ListCategoryResp, 0)
	for _, v := range categories {
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			v.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + v.ImgPath
		}

		cResp = append(cResp, &types.ListCategoryResp{
			ID:           v.ID,
			CategoryName: v.CategoryName,
			CreatedAt:    v.CreatedAt.Unix(),
			ImgPath: v.ImgPath,
		})
	}

	resp = &types.DataListResp{
		Item:  cResp,
		Total: int64(len(cResp)),
	}

	return
}
