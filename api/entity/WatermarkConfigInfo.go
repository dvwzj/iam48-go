package entity

type WatermarkConfigInfo struct {
	// @SerializedName("isAllowOnScreen")
	// @Nullable
	// private Boolean isAllowOnScreen;
	IsAllowOnScreen *bool `json:"isAllowOnScreen"`
	// @SerializedName("isDisplayWatermark")
	// @Nullable
	// private Boolean isDisplayWatermark;
	IsDisplayWatermark *bool `json:"isDisplayWatermark"`
	// @SerializedName("watermarkDisplayAfterSeconds")
	// @Nullable
	// private Integer watermarkDisplayAfterSeconds;
	WatermarkDisplayAfterSeconds *int `json:"watermarkDisplayAfterSeconds"`
	// @SerializedName("watermarkDisplaySeconds")
	// @Nullable
	// private Integer watermarkDisplaySeconds;
	WatermarkDisplaySeconds *int `json:"watermarkDisplaySeconds"`
	// @SerializedName("watermarkIntervalBaseSeconds")
	// @Nullable
	// private Integer watermarkIntervalBaseSeconds;
	WatermarkIntervalBaseSeconds *int `json:"watermarkIntervalBaseSeconds"`
	// @SerializedName("watermarkIntervalRandomSecondsFrom")
	// @Nullable
	// private Integer watermarkIntervalRandomSecondsFrom;
	WatermarkIntervalRandomSecondsFrom *int `json:"watermarkIntervalRandomSecondsFrom"`
	// @SerializedName("watermarkIntervalRandomSecondsTo")
	// @Nullable
	// private Integer watermarkIntervalRandomSecondsTo;
	WatermarkIntervalRandomSecondsTo *int `json:"watermarkIntervalRandomSecondsTo"`
}
