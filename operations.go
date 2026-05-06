package wdt

var (
	OpListProduct = Operation{
		Name:         "ListProduct",
		DirectMethod: "goods.Goods.queryWithSpec",
		QimenMethod:  "wdt.goods.goods.querywithspec",
	}
	OpGetProduct = Operation{
		Name:         "GetProduct",
		DirectMethod: "goods.Goods.queryWithSpec",
		QimenMethod:  "wdt.goods.goods.querywithspec",
	}
	OpCreateProduct = Operation{Name: "CreateProduct", DirectMethod: "goods.Goods.push"}
	OpListCategory  = Operation{Name: "ListCategory", DirectMethod: "goods.GoodsClass.search", QimenMethod: "wdt.goods.class.query"}

	OpListSupplier          = Operation{Name: "ListSupplier", DirectMethod: "setting.PurchaseProvider.queryDetail", QimenMethod: "wdt.purchase.provider.query"}
	OpCreateSupplier        = Operation{Name: "CreateSupplier", DirectMethod: "setting.PurchaseProvider.push"}
	OpCreateProductSupplier = Operation{Name: "CreateProductSupplier", DirectMethod: "purchase.ProviderGoods.upload", DirectBody: DirectBodyArray}

	OpListWarehouse        = Operation{Name: "ListWarehouse", DirectMethod: "setting.Warehouse.queryWarehouse", QimenMethod: "wdt.warehouse.query"}
	OpListShop             = Operation{Name: "ListShop", DirectMethod: "setting.Shop.queryShop", QimenMethod: "wdt.shop.query"}
	OpListLogisticsCompany = Operation{Name: "ListLogisticsCompany", DirectMethod: "setting.Logistics.queryLogistics", QimenMethod: "wdt.setting.logistics.querylogistics"}

	OpCreateRawOrder      = Operation{Name: "CreateRawOrder", DirectMethod: "sales.Trade.push", DirectBody: DirectBodyArray}
	OpListRawOrder        = Operation{Name: "ListRawOrder", QimenMethod: "wdt.sales.rawtrade.search"}
	OpListHistoryRawOrder = Operation{Name: "ListHistoryRawOrder", QimenMethod: "wdt.sales.rawtrade.searchhistory"}

	OpListSaleOrder         = Operation{Name: "ListSaleOrder", QimenMethod: "wdt.sales.tradequery.querywithdetail"}
	OpListArchivedSaleOrder = Operation{Name: "ListArchivedSaleOrder", QimenMethod: "wdt.sales.tradequery.queryhistorywithdetail"}
	OpListReturnOrder       = Operation{Name: "ListReturnOrder", QimenMethod: "wdt.aftersales.refund.refund.search"}
	OpListArchivedReturn    = Operation{Name: "ListArchivedReturnOrder", QimenMethod: "wdt.aftersales.refund.refund.searchhistory"}
	OpListSaleStockout      = Operation{Name: "ListSaleStockout", QimenMethod: "wdt.wms.stockout.sales.querywithdetail"}
	OpListSalesPayment      = Operation{Name: "ListSalesPayment", QimenMethod: "wdt.sales.payment.querywithdetail"}

	OpCreatePurchaseOrder           = Operation{Name: "CreatePurchaseOrder", DirectMethod: "purchase.PurchaseOrder.createOrder"}
	OpListPurchaseOrder             = Operation{Name: "ListPurchaseOrder", DirectMethod: "purchase.PurchaseOrder.queryWithDetail", QimenMethod: "wdt.purchase.purchaseorder.querywithdetail"}
	OpGetPurchaseOrder              = Operation{Name: "GetPurchaseOrder", DirectMethod: "purchase.PurchaseOrder.queryWithDetail", QimenMethod: "wdt.purchase.purchaseorder.querywithdetail"}
	OpCancelOrStopPurchaseOrder     = Operation{Name: "CancelOrStopPurchaseOrder", DirectMethod: "purchase.PurchaseOrder.cancelByType"}
	OpListPurchaseStockinOrder      = Operation{Name: "ListPurchaseStockinOrder", DirectMethod: "wms.stockin.Purchase.queryWithDetail"}
	OpCreatePurchaseReturn          = Operation{Name: "CreatePurchaseReturn", DirectMethod: "purchase.PurchaseReturn.createOrder", DirectBody: DirectBodyArray}
	OpListPurchaseReturn            = Operation{Name: "ListPurchaseReturn", DirectMethod: "purchase.PurchaseReturn.queryWithDetail", QimenMethod: "wdt.purchase.purchasereturn.querywithdetail"}
	OpCancelPurchaseReturn          = Operation{Name: "CancelPurchaseReturn", DirectMethod: "purchase.PurchaseReturn.cancelOrder"}
	OpStopPurchaseReturnStockout    = Operation{Name: "StopPurchaseReturnStockout", DirectMethod: "purchase.PurchaseReturn.pending"}
	OpCreateStockTransferByPosition = Operation{Name: "CreateStockTransferByPositionNo", DirectMethod: "wms.stocktransfer.Edit.createOrder", DirectBody: DirectBodyArray}

	OpGetStockinDoc        = Operation{Name: "GetStockinDoc", DirectMethod: "wms.stockin.Purchase.queryWithDetail"}
	OpCreateStockinRecord  = Operation{Name: "CreateStockinRecord", DirectMethod: "wms.stockin.Purchase.upload", DirectBody: DirectBodyArray}
	OpListPrestockin       = Operation{Name: "ListPrestockin", DirectMethod: "wms.stockin.PreStockin.search", QimenMethod: "wdt.wms.stockin.prestockin.search"}
	OpListReturnStockin    = Operation{Name: "ListReturnStockin", QimenMethod: "wdt.wms.stockin.refund.querywithdetail"}
	OpListReturnPrestockin = Operation{Name: "ListReturnPrestockin", QimenMethod: "wdt.wms.stockin.prestockin.search"}

	OpListStockByPage      = Operation{Name: "ListStockByPage", DirectMethod: "wms.StockSpec.search2", QimenMethod: "wdt.wms.stockspec.search"}
	OpListStock            = Operation{Name: "ListStock", DirectMethod: "wms.StockSpec.search", QimenMethod: "wdt.wms.stockspec.search"}
	OpGetStockDetail       = Operation{Name: "GetStockDetail", DirectMethod: "wms.StockSpec.stockDetailSearch"}
	OpListDefaultPosition  = Operation{Name: "ListDefaultPosition", DirectMethod: "wms.PositionCapacity.search"}
	OpListPositionStock    = Operation{Name: "ListPositionStock", DirectMethod: "wms.PositionCapacity.search"}
	OpListVirtualStock     = Operation{Name: "ListVirtualStock", DirectMethod: "setting.strategy.VirtualWarehouse.stockSearch"}
	OpListVirtualWarehouse = Operation{
		Name:         "ListVirtualWarehouse",
		DirectMethod: "setting.strategy.VirtualWarehouse.warehouseSearch",
	}
	OpListStockDefect = Operation{Name: "ListStockDefect", QimenMethod: "wdt.wms.stockdefect.defectchange.search"}
	OpListStockSpec   = Operation{Name: "ListStockSpec", QimenMethod: "wdt.wms.stockspec.search"}

	OpListLogistics       = Operation{Name: "ListLogistics", DirectMethod: "wms.stockout.Sales.searchLogistics"}
	OpListOtherStockout   = Operation{Name: "ListOtherStockout", QimenMethod: "wdt.wms.stockout.otherquery.querywithdetail"}
	OpListPlatformProduct = Operation{Name: "ListPlatformProduct", QimenMethod: "wdt.goods.apigoods.search"}
	OpListSku             = Operation{Name: "ListSku", QimenMethod: "wdt.goods.suite.search"}
	OpReissueOrder        = Operation{Name: "ReissueOrder", DirectMethod: "sales.TradeManual.reissueOrder", DirectBody: DirectBodyArray}
)

var Operations = []Operation{
	OpListProduct, OpGetProduct, OpCreateProduct, OpListCategory,
	OpListSupplier, OpCreateSupplier, OpCreateProductSupplier,
	OpListWarehouse, OpListShop, OpListLogisticsCompany,
	OpCreateRawOrder, OpListRawOrder, OpListHistoryRawOrder,
	OpListSaleOrder, OpListArchivedSaleOrder, OpListReturnOrder, OpListArchivedReturn,
	OpListSaleStockout, OpListSalesPayment,
	OpCreatePurchaseOrder, OpListPurchaseOrder, OpGetPurchaseOrder, OpCancelOrStopPurchaseOrder,
	OpListPurchaseStockinOrder, OpCreatePurchaseReturn, OpListPurchaseReturn, OpCancelPurchaseReturn, OpStopPurchaseReturnStockout,
	OpCreateStockTransferByPosition, OpGetStockinDoc, OpCreateStockinRecord, OpListPrestockin,
	OpListReturnStockin, OpListReturnPrestockin,
	OpListStockByPage, OpListStock, OpGetStockDetail, OpListDefaultPosition, OpListPositionStock,
	OpListVirtualStock, OpListVirtualWarehouse, OpListStockDefect, OpListStockSpec,
	OpListLogistics, OpListOtherStockout, OpListPlatformProduct, OpListSku, OpReissueOrder,
}
