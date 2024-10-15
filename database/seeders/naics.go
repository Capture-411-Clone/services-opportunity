package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func NaicsesSeeder(db *gorm.DB) {
	Naicses := []models.Naics{

		{
			UserID: 1,
			Name:   "111110 Soybean Farming",
		},

		{
			UserID: 1,
			Name:   "111120 Oilseed (except Soybean) Farming",
		},

		{
			UserID: 1,
			Name:   "111130 Dry Pea and Bean Farming",
		},

		{
			UserID: 1,
			Name:   "111140 Wheat Farming",
		},

		{
			UserID: 1,
			Name:   "111150 Corn Farming",
		},

		{
			UserID: 1,
			Name:   "111160 Rice Farming",
		},

		{
			UserID: 1,
			Name:   "111191 Oilseed and Grain Combination Farming",
		},

		{
			UserID: 1,
			Name:   "111199 All Other Grain Farming",
		},

		{
			UserID: 1,
			Name:   "111211 Potato Farming",
		},

		{
			UserID: 1,
			Name:   "111219 Other Vegetable (except Potato) and Melon Farming",
		},

		{
			UserID: 1,
			Name:   "111310 Orange Groves",
		},

		{
			UserID: 1,
			Name:   "111320 Citrus (except Orange) Groves",
		},

		{
			UserID: 1,
			Name:   "111331 Apple Orchards",
		},

		{
			UserID: 1,
			Name:   "111332 Grape Vineyards",
		},

		{
			UserID: 1,
			Name:   "111333 Strawberry Farming",
		},

		{
			UserID: 1,
			Name:   "111334 Berry (except Strawberry) Farming",
		},

		{
			UserID: 1,
			Name:   "111335 Tree Nut Farming",
		},

		{
			UserID: 1,
			Name:   "111336 Fruit and Tree Nut Combination Farming",
		},

		{
			UserID: 1,
			Name:   "111339 Other Noncitrus Fruit Farming",
		},

		{
			UserID: 1,
			Name:   "111411 Mushroom Production",
		},

		{
			UserID: 1,
			Name:   "111419 Other Food Crops Grown Under Cover",
		},

		{
			UserID: 1,
			Name:   "111421 Nursery and Tree Production",
		},

		{
			UserID: 1,
			Name:   "111422 Floriculture Production",
		},

		{
			UserID: 1,
			Name:   "111910 Tobacco Farming",
		},

		{
			UserID: 1,
			Name:   "111920 Cotton Farming",
		},

		{
			UserID: 1,
			Name:   "111930 Sugarcane Farming",
		},

		{
			UserID: 1,
			Name:   "111940 Hay Farming",
		},

		{
			UserID: 1,
			Name:   "111991 Sugar Beet Farming",
		},

		{
			UserID: 1,
			Name:   "111992 Peanut Farming",
		},

		{
			UserID: 1,
			Name:   "111998 All Other Miscellaneous Crop Farming",
		},

		{
			UserID: 1,
			Name:   "112111 Beef Cattle Ranching and Farming",
		},

		{
			UserID: 1,
			Name:   "112112 Cattle Feedlots",
		},

		{
			UserID: 1,
			Name:   "112120 Dairy Cattle and Milk Production",
		},

		{
			UserID: 1,
			Name:   "112130 Dual-Purpose Cattle Ranching and Farming",
		},

		{
			UserID: 1,
			Name:   "112210 Hog and Pig Farming",
		},

		{
			UserID: 1,
			Name:   "112310 Chicken Egg Production",
		},

		{
			UserID: 1,
			Name:   "112320 Broilers and Other Meat Type Chicken Production",
		},

		{
			UserID: 1,
			Name:   "112330 Turkey Production",
		},

		{
			UserID: 1,
			Name:   "112340 Poultry Hatcheries",
		},

		{
			UserID: 1,
			Name:   "112390 Other Poultry Production",
		},

		{
			UserID: 1,
			Name:   "112410 Sheep Farming",
		},

		{
			UserID: 1,
			Name:   "112420 Goat Farming",
		},

		{
			UserID: 1,
			Name:   "112511 Finfish Farming and Fish Hatcheries",
		},

		{
			UserID: 1,
			Name:   "112512 Shellfish Farming",
		},

		{
			UserID: 1,
			Name:   "112519 Other Aquaculture",
		},

		{
			UserID: 1,
			Name:   "112910 Apiculture",
		},

		{
			UserID: 1,
			Name:   "112920 Horses and Other Equine Production",
		},

		{
			UserID: 1,
			Name:   "112930 Fur-Bearing Animal and Rabbit Production",
		},

		{
			UserID: 1,
			Name:   "112990 All Other Animal Production",
		},

		{
			UserID: 1,
			Name:   "113110 Timber Tract Operations",
		},

		{
			UserID: 1,
			Name:   "113210 Forest Nurseries and Gathering of Forest Products",
		},

		{
			UserID: 1,
			Name:   "113310 Logging",
		},

		{
			UserID: 1,
			Name:   "114111 Finfish Fishing",
		},

		{
			UserID: 1,
			Name:   "114112 Shellfish Fishing",
		},

		{
			UserID: 1,
			Name:   "114119 Other Marine Fishing",
		},

		{
			UserID: 1,
			Name:   "114210 Hunting and Trapping",
		},

		{
			UserID: 1,
			Name:   "115111 Cotton Ginning",
		},

		{
			UserID: 1,
			Name:   "115112 Soil Preparation, Planting, and Cultivating",
		},

		{
			UserID: 1,
			Name:   "115113 Crop Harvesting, Primarily by Machine",
		},

		{
			UserID: 1,
			Name:   "115114 Postharvest Crop Activities (except Cotton Ginning)",
		},

		{
			UserID: 1,
			Name:   "115115 Farm Labor Contractors and Crew Leaders",
		},

		{
			UserID: 1,
			Name:   "115116 Farm Management Services",
		},

		{
			UserID: 1,
			Name:   "115210 Support Activities for Animal Production",
		},

		{
			UserID: 1,
			Name:   "115310 Support Activities for Forestry",
		},

		{
			UserID: 1,
			Name:   "211120 Crude Petroleum Extraction",
		},

		{
			UserID: 1,
			Name:   "211130 Natural Gas Extraction",
		},

		{
			UserID: 1,
			Name:   "212114 Surface Coal Mining",
		},

		{
			UserID: 1,
			Name:   "212115 Underground Coal Mining",
		},

		{
			UserID: 1,
			Name:   "212210 Iron Ore Mining",
		},

		{
			UserID: 1,
			Name:   "212220 Gold Ore and Silver Ore Mining",
		},

		{
			UserID: 1,
			Name:   "212230 Copper, Nickel, Lead, and Zinc Mining",
		},

		{
			UserID: 1,
			Name:   "212290 Other Metal Ore Mining",
		},

		{
			UserID: 1,
			Name:   "212311 Dimension Stone Mining and Quarrying",
		},

		{
			UserID: 1,
			Name:   "212312 Crushed and Broken Limestone Mining and Quarrying",
		},

		{
			UserID: 1,
			Name:   "212313 Crushed and Broken Granite Mining and Quarrying",
		},

		{
			UserID: 1,
			Name:   "212319 Other Crushed and Broken Stone Mining and Quarrying",
		},

		{
			UserID: 1,
			Name:   "212321 Construction Sand and Gravel Mining",
		},

		{
			UserID: 1,
			Name:   "212322 Industrial Sand Mining",
		},

		{
			UserID: 1,
			Name:   "212323 Kaolin, Clay, and Ceramic and Refractory Minerals Mining",
		},

		{
			UserID: 1,
			Name:   "212390 Other Nonmetallic Mineral Mining and Quarrying",
		},

		{
			UserID: 1,
			Name:   "213111 Drilling Oil and Gas Wells",
		},

		{
			UserID: 1,
			Name:   "213112 Support Activities for Oil and Gas Operations",
		},

		{
			UserID: 1,
			Name:   "213113 Support Activities for Coal Mining",
		},

		{
			UserID: 1,
			Name:   "213114 Support Activities for Metal Mining",
		},

		{
			UserID: 1,
			Name:   "213115 Support Activities for Nonmetallic Minerals (except Fuels) Mining",
		},

		{
			UserID: 1,
			Name:   "221111 Hydroelectric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221112 Fossil Fuel Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221113 Nuclear Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221114 Solar Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221115 Wind Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221116 Geothermal Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221117 Biomass Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221118 Other Electric Power Generation",
		},

		{
			UserID: 1,
			Name:   "221121 Electric Bulk Power Transmission and Control",
		},

		{
			UserID: 1,
			Name:   "221122 Electric Power Distribution",
		},

		{
			UserID: 1,
			Name:   "221210 Natural Gas Distribution",
		},

		{
			UserID: 1,
			Name:   "221310 Water Supply and Irrigation Systems",
		},

		{
			UserID: 1,
			Name:   "221320 Sewage Treatment Facilities",
		},

		{
			UserID: 1,
			Name:   "221330 Steam and Air-Conditioning Supply",
		},

		{
			UserID: 1,
			Name:   "236115 New Single-Family Housing Construction (except For-Sale Builders)",
		},

		{
			UserID: 1,
			Name:   "236116 New Multifamily Housing Construction (except For-Sale Builders)",
		},

		{
			UserID: 1,
			Name:   "236117 New Housing For-Sale Builders",
		},

		{
			UserID: 1,
			Name:   "236118 Residential Remodelers",
		},

		{
			UserID: 1,
			Name:   "236210 Industrial Building Construction",
		},

		{
			UserID: 1,
			Name:   "236220 Commercial and Institutional Building Construction",
		},

		{
			UserID: 1,
			Name:   "237110 Water and Sewer Line and Related Structures Construction",
		},

		{
			UserID: 1,
			Name:   "237120 Oil and Gas Pipeline and Related Structures Construction",
		},

		{
			UserID: 1,
			Name:   "237130 Power and Communication Line and Related Structures Construction",
		},

		{
			UserID: 1,
			Name:   "237210 Land Subdivision",
		},

		{
			UserID: 1,
			Name:   "237310 Highway, Street, and Bridge Construction",
		},

		{
			UserID: 1,
			Name:   "237990 Other Heavy and Civil Engineering Construction",
		},

		{
			UserID: 1,
			Name:   "238110 Poured Concrete Foundation and Structure Contractors",
		},

		{
			UserID: 1,
			Name:   "238120 Structural Steel and Precast Concrete Contractors",
		},

		{
			UserID: 1,
			Name:   "238130 Framing Contractors",
		},

		{
			UserID: 1,
			Name:   "238140 Masonry Contractors",
		},

		{
			UserID: 1,
			Name:   "238150 Glass and Glazing Contractors",
		},

		{
			UserID: 1,
			Name:   "238160 Roofing Contractors",
		},

		{
			UserID: 1,
			Name:   "238170 Siding Contractors",
		},

		{
			UserID: 1,
			Name:   "238190 Other Foundation, Structure, and Building Exterior Contractors",
		},

		{
			UserID: 1,
			Name:   "238210 Electrical Contractors and Other Wiring Installation Contractors",
		},

		{
			UserID: 1,
			Name:   "238220 Plumbing, Heating, and Air-Conditioning Contractors",
		},

		{
			UserID: 1,
			Name:   "238290 Other Building Equipment Contractors",
		},

		{
			UserID: 1,
			Name:   "238310 Drywall and Insulation Contractors",
		},

		{
			UserID: 1,
			Name:   "238320 Painting and Wall Covering Contractors",
		},

		{
			UserID: 1,
			Name:   "238330 Flooring Contractors",
		},

		{
			UserID: 1,
			Name:   "238340 Tile and Terrazzo Contractors",
		},

		{
			UserID: 1,
			Name:   "238350 Finish Carpentry Contractors",
		},

		{
			UserID: 1,
			Name:   "238390 Other Building Finishing Contractors",
		},

		{
			UserID: 1,
			Name:   "238910 Site Preparation Contractors",
		},

		{
			UserID: 1,
			Name:   "238990 All Other Specialty Trade Contractors",
		},

		{
			UserID: 1,
			Name:   "311111 Dog and Cat Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311119 Other Animal Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311211 Flour Milling",
		},

		{
			UserID: 1,
			Name:   "311212 Rice Milling",
		},

		{
			UserID: 1,
			Name:   "311213 Malt Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311221 Wet Corn Milling and Starch Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311224 Soybean and Other Oilseed Processing",
		},

		{
			UserID: 1,
			Name:   "311225 Fats and Oils Refining and Blending",
		},

		{
			UserID: 1,
			Name:   "311230 Breakfast Cereal Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311313 Beet Sugar Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311314 Cane Sugar Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311340 Nonchocolate Confectionery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311351 Chocolate and Confectionery Manufacturing from Cacao Beans",
		},

		{
			UserID: 1,
			Name:   "311352 Confectionery Manufacturing from Purchased Chocolate",
		},

		{
			UserID: 1,
			Name:   "311411 Frozen Fruit, Juice, and Vegetable Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311412 Frozen Specialty Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311421 Fruit and Vegetable Canning",
		},

		{
			UserID: 1,
			Name:   "311422 Specialty Canning",
		},

		{
			UserID: 1,
			Name:   "311423 Dried and Dehydrated Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311511 Fluid Milk Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311512 Creamery Butter Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311513 Cheese Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311514 Dry, Condensed, and Evaporated Dairy Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311520 Ice Cream and Frozen Dessert Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311611 Animal (except Poultry) Slaughtering",
		},

		{
			UserID: 1,
			Name:   "311612 Meat Processed from Carcasses",
		},

		{
			UserID: 1,
			Name:   "311613 Rendering and Meat Byproduct Processing",
		},

		{
			UserID: 1,
			Name:   "311615 Poultry Processing",
		},

		{
			UserID: 1,
			Name:   "311710 Seafood Product Preparation and Packaging",
		},

		{
			UserID: 1,
			Name:   "311811 Retail Bakeries",
		},

		{
			UserID: 1,
			Name:   "311812 Commercial Bakeries",
		},

		{
			UserID: 1,
			Name:   "311813 Frozen Cakes, Pies, and Other Pastries Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311821 Cookie and Cracker Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311824 Dry Pasta, Dough, and Flour Mixes Manufacturing from Purchased Flour",
		},

		{
			UserID: 1,
			Name:   "311830 Tortilla Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311911 Roasted Nuts and Peanut Butter Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311919 Other Snack Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311920 Coffee and Tea Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311930 Flavoring Syrup and Concentrate Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311941 Mayonnaise, Dressing, and Other Prepared Sauce Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311942 Spice and Extract Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311991 Perishable Prepared Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "311999 All Other Miscellaneous Food Manufacturing",
		},

		{
			UserID: 1,
			Name:   "312111 Soft Drink Manufacturing",
		},

		{
			UserID: 1,
			Name:   "312112 Bottled Water Manufacturing",
		},

		{
			UserID: 1,
			Name:   "312113 Ice Manufacturing",
		},

		{
			UserID: 1,
			Name:   "312120 Breweries",
		},

		{
			UserID: 1,
			Name:   "312130 Wineries",
		},

		{
			UserID: 1,
			Name:   "312140 Distilleries",
		},

		{
			UserID: 1,
			Name:   "312230 Tobacco Manufacturing",
		},

		{
			UserID: 1,
			Name:   "313110 Fiber, Yarn, and Thread Mills",
		},

		{
			UserID: 1,
			Name:   "313210 Broadwoven Fabric Mills",
		},

		{
			UserID: 1,
			Name:   "313220 Narrow Fabric Mills and Schiffli Machine Embroidery",
		},

		{
			UserID: 1,
			Name:   "313230 Nonwoven Fabric Mills",
		},

		{
			UserID: 1,
			Name:   "313240 Knit Fabric Mills",
		},

		{
			UserID: 1,
			Name:   "313310 Textile and Fabric Finishing Mills",
		},

		{
			UserID: 1,
			Name:   "313320 Fabric Coating Mills",
		},

		{
			UserID: 1,
			Name:   "314110 Carpet and Rug Mills",
		},

		{
			UserID: 1,
			Name:   "314120 Curtain and Linen Mills",
		},

		{
			UserID: 1,
			Name:   "314910 Textile Bag and Canvas Mills",
		},

		{
			UserID: 1,
			Name:   "314994 Rope, Cordage, Twine, Tire Cord, and Tire Fabric Mills",
		},

		{
			UserID: 1,
			Name:   "314999 All Other Miscellaneous Textile Product Mills",
		},

		{
			UserID: 1,
			Name:   "315120 Apparel Knitting Mills",
		},

		{
			UserID: 1,
			Name:   "315210 Cut and Sew Apparel Contractors",
		},

		{
			UserID: 1,
			Name:   "315250 Cut and Sew Apparel Manufacturing (except Contractors)",
		},

		{
			UserID: 1,
			Name:   "315990 Apparel Accessories and Other Apparel Manufacturing",
		},

		{
			UserID: 1,
			Name:   "316110 Leather and Hide Tanning and Finishing",
		},

		{
			UserID: 1,
			Name:   "316210 Footwear Manufacturing",
		},

		{
			UserID: 1,
			Name:   "316990 Other Leather and Allied Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321113 Sawmills",
		},

		{
			UserID: 1,
			Name:   "321114 Wood Preservation",
		},

		{
			UserID: 1,
			Name:   "321211 Hardwood Veneer and Plywood Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321212 Softwood Veneer and Plywood Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321215 Engineered Wood Member Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321219 Reconstituted Wood Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321911 Wood Window and Door Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321912 Cut Stock, Resawing Lumber, and Planing",
		},

		{
			UserID: 1,
			Name:   "321918 Other Millwork (including Flooring)",
		},

		{
			UserID: 1,
			Name:   "321920 Wood Container and Pallet Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321991 Manufactured Home (Mobile Home) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321992 Prefabricated Wood Building Manufacturing",
		},

		{
			UserID: 1,
			Name:   "321999 All Other Miscellaneous Wood Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322110 Pulp Mills",
		},

		{
			UserID: 1,
			Name:   "322120 Paper Mills",
		},

		{
			UserID: 1,
			Name:   "322130 Paperboard Mills",
		},

		{
			UserID: 1,
			Name:   "322211 Corrugated and Solid Fiber Box Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322212 Folding Paperboard Box Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322219 Other Paperboard Container Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322220 Paper Bag and Coated and Treated Paper Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322230 Stationery Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322291 Sanitary Paper Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "322299 All Other Converted Paper Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "323111 Commercial Printing (except Screen and Books)",
		},

		{
			UserID: 1,
			Name:   "323113 Commercial Screen Printing",
		},

		{
			UserID: 1,
			Name:   "323117 Books Printing",
		},

		{
			UserID: 1,
			Name:   "323120 Support Activities for Printing",
		},

		{
			UserID: 1,
			Name:   "324110 Petroleum Refineries",
		},

		{
			UserID: 1,
			Name:   "324121 Asphalt Paving Mixture and Block Manufacturing",
		},

		{
			UserID: 1,
			Name:   "324122 Asphalt Shingle and Coating Materials Manufacturing",
		},

		{
			UserID: 1,
			Name:   "324191 Petroleum Lubricating Oil and Grease Manufacturing",
		},

		{
			UserID: 1,
			Name:   "324199 All Other Petroleum and Coal Products Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325110 Petrochemical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325120 Industrial Gas Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325130 Synthetic Dye and Pigment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325180 Other Basic Inorganic Chemical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325193 Ethyl Alcohol Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325194 Cyclic Crude, Intermediate, and Gum and Wood Chemical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325199 All Other Basic Organic Chemical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325211 Plastics Material and Resin Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325212 Synthetic Rubber Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325220 Artificial and Synthetic Fibers and Filaments Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325311 Nitrogenous Fertilizer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325312 Phosphatic Fertilizer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325314 Fertilizer (Mixing Only) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325315 Compost Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325320 Pesticide and Other Agricultural Chemical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325411 Medicinal and Botanical Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325412 Pharmaceutical Preparation Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325413 In-Vitro Diagnostic Substance Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325414 Biological Product (except Diagnostic) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325510 Paint and Coating Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325520 Adhesive Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325611 Soap and Other Detergent Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325612 Polish and Other Sanitation Good Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325613 Surface Active Agent Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325620 Toilet Preparation Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325910 Printing Ink Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325920 Explosives Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325991 Custom Compounding of Purchased Resins",
		},

		{
			UserID: 1,
			Name:   "325992 Photographic Film, Paper, Plate, Chemical, and Copy Toner Manufacturing",
		},

		{
			UserID: 1,
			Name:   "325998 All Other Miscellaneous Chemical Product and Preparation Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326111 Plastics Bag and Pouch Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326112 Plastics Packaging Film and Sheet (including Laminated) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326113 Unlaminated Plastics Film and Sheet (except Packaging) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326121 Unlaminated Plastics Profile Shape Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326122 Plastics Pipe and Pipe Fitting Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326130 Laminated Plastics Plate, Sheet (except Packaging), and Shape Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326140 Polystyrene Foam Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326150 Urethane and Other Foam Product (except Polystyrene) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326160 Plastics Bottle Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326191 Plastics Plumbing Fixture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326199 All Other Plastics Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326211 Tire Manufacturing (except Retreading)",
		},

		{
			UserID: 1,
			Name:   "326212 Tire Retreading",
		},

		{
			UserID: 1,
			Name:   "326220 Rubber and Plastics Hoses and Belting Manufacturing",
		},

		{
			UserID: 1,
			Name:   "326291 Rubber Product Manufacturing for Mechanical Use",
		},

		{
			UserID: 1,
			Name:   "326299 All Other Rubber Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327110 Pottery, Ceramics, and Plumbing Fixture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327120 Clay Building Material and Refractories Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327211 Flat Glass Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327212 Other Pressed and Blown Glass and Glassware Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327213 Glass Container Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327215 Glass Product Manufacturing Made of Purchased Glass",
		},

		{
			UserID: 1,
			Name:   "327310 Cement Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327320 Ready-Mix Concrete Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327331 Concrete Block and Brick Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327332 Concrete Pipe Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327390 Other Concrete Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327410 Lime Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327420 Gypsum Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327910 Abrasive Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327991 Cut Stone and Stone Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327992 Ground or Treated Mineral and Earth Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327993 Mineral Wool Manufacturing",
		},

		{
			UserID: 1,
			Name:   "327999 All Other Miscellaneous Nonmetallic Mineral Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "331110 Iron and Steel Mills and Ferroalloy Manufacturing",
		},

		{
			UserID: 1,
			Name:   "331210 Iron and Steel Pipe and Tube Manufacturing from Purchased Steel",
		},

		{
			UserID: 1,
			Name:   "331221 Rolled Steel Shape Manufacturing",
		},

		{
			UserID: 1,
			Name:   "331222 Steel Wire Drawing",
		},

		{
			UserID: 1,
			Name:   "331313 Alumina Refining and Primary Aluminum Production",
		},

		{
			UserID: 1,
			Name:   "331314 Secondary Smelting and Alloying of Aluminum",
		},

		{
			UserID: 1,
			Name:   "331315 Aluminum Sheet, Plate, and Foil Manufacturing",
		},

		{
			UserID: 1,
			Name:   "331318 Other Aluminum Rolling, Drawing, and Extruding",
		},

		{
			UserID: 1,
			Name:   "331410 Nonferrous Metal (except Aluminum) Smelting and Refining",
		},

		{
			UserID: 1,
			Name:   "331420 Copper Rolling, Drawing, Extruding, and Alloying",
		},

		{
			UserID: 1,
			Name:   "331491 Nonferrous Metal (except Copper and Aluminum) Rolling, Drawing, and Extruding",
		},

		{
			UserID: 1,
			Name:   "331492 Secondary Smelting, Refining, and Alloying of Nonferrous Metal (except Copper and Aluminum)",
		},

		{
			UserID: 1,
			Name:   "331511 Iron Foundries",
		},

		{
			UserID: 1,
			Name:   "331512 Steel Investment Foundries",
		},

		{
			UserID: 1,
			Name:   "331513 Steel Foundries (except Investment)",
		},

		{
			UserID: 1,
			Name:   "331523 Nonferrous Metal Die-Casting Foundries",
		},

		{
			UserID: 1,
			Name:   "331524 Aluminum Foundries (except Die-Casting)",
		},

		{
			UserID: 1,
			Name:   "331529 Other Nonferrous Metal Foundries (except Die-Casting)",
		},

		{
			UserID: 1,
			Name:   "332111 Iron and Steel Forging",
		},

		{
			UserID: 1,
			Name:   "332112 Nonferrous Forging",
		},

		{
			UserID: 1,
			Name:   "332114 Custom Roll Forming",
		},

		{
			UserID: 1,
			Name:   "332117 Powder Metallurgy Part Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332119 Metal Crown, Closure, and Other Metal Stamping (except Automotive)",
		},

		{
			UserID: 1,
			Name:   "332215 Metal Kitchen Cookware, Utensil, Cutlery, and Flatware (except Precious) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332216 Saw Blade and Handtool Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332311 Prefabricated Metal Building and Component Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332312 Fabricated Structural Metal Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332313 Plate Work Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332321 Metal Window and Door Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332322 Sheet Metal Work Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332323 Ornamental and Architectural Metal Work Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332410 Power Boiler and Heat Exchanger Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332420 Metal Tank (Heavy Gauge) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332431 Metal Can Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332439 Other Metal Container Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332510 Hardware Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332613 Spring Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332618 Other Fabricated Wire Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332710 Machine Shops",
		},

		{
			UserID: 1,
			Name:   "332721 Precision Turned Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332722 Bolt, Nut, Screw, Rivet, and Washer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332811 Metal Heat Treating",
		},

		{
			UserID: 1,
			Name:   "332812 Metal Coating, Engraving (except Jewelry and Silverware), and Allied Services to Manufacturers",
		},

		{
			UserID: 1,
			Name:   "332813 Electroplating, Plating, Polishing, Anodizing, and Coloring",
		},

		{
			UserID: 1,
			Name:   "332911 Industrial Valve Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332912 Fluid Power Valve and Hose Fitting Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332913 Plumbing Fixture Fitting and Trim Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332919 Other Metal Valve and Pipe Fitting Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332991 Ball and Roller Bearing Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332992 Small Arms Ammunition Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332993 Ammunition (except Small Arms) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332994 Small Arms, Ordnance, and Ordnance Accessories Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332996 Fabricated Pipe and Pipe Fitting Manufacturing",
		},

		{
			UserID: 1,
			Name:   "332999 All Other Miscellaneous Fabricated Metal Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333111 Farm Machinery and Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333112 Lawn and Garden Tractor and Home Lawn and Garden Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333120 Construction Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333131 Mining Machinery and Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333132 Oil and Gas Field Machinery and Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333241 Food Product Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333242 Semiconductor Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333243 Sawmill, Woodworking, and Paper Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333248 All Other Industrial Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333310 Commercial and Service Industry Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333413 Industrial and Commercial Fan and Blower and Air Purification Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333414 Heating Equipment (except Warm Air Furnaces) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333415 Air-Conditioning and Warm Air Heating Equipment and Commercial and Industrial Refrigeration Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333511 Industrial Mold Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333514 Special Die and Tool, Die Set, Jig, and Fixture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333515 Cutting Tool and Machine Tool Accessory Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333517 Machine Tool Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333519 Rolling Mill and Other Metalworking Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333611 Turbine and Turbine Generator Set Units Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333612 Speed Changer, Industrial High-Speed Drive, and Gear Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333613 Mechanical Power Transmission Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333618 Other Engine Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333912 Air and Gas Compressor Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333914 Measuring, Dispensing, and Other Pumping Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333921 Elevator and Moving Stairway Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333922 Conveyor and Conveying Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333923 Overhead Traveling Crane, Hoist, and Monorail System Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333924 Industrial Truck, Tractor, Trailer, and Stacker Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333991 Power-Driven Handtool Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333992 Welding and Soldering Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333993 Packaging Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333994 Industrial Process Furnace and Oven Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333995 Fluid Power Cylinder and Actuator Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333996 Fluid Power Pump and Motor Manufacturing",
		},

		{
			UserID: 1,
			Name:   "333998 All Other Miscellaneous General Purpose Machinery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334111 Electronic Computer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334112 Computer Storage Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334118 Computer Terminal and Other Computer Peripheral Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334210 Telephone Apparatus Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334220 Radio and Television Broadcasting and Wireless Communications Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334290 Other Communications Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334310 Audio and Video Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334412 Bare Printed Circuit Board Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334413 Semiconductor and Related Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334416 Capacitor, Resistor, Coil, Transformer, and Other Inductor Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334417 Electronic Connector Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334418 Printed Circuit Assembly (Electronic Assembly) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334419 Other Electronic Component Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334510 Electromedical and Electrotherapeutic Apparatus Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334511 Search, Detection, Navigation, Guidance, Aeronautical, and Nautical System and Instrument Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334512 Automatic Environmental Control Manufacturing for Residential, Commercial, and Appliance Use",
		},

		{
			UserID: 1,
			Name:   "334513 Instruments and Related Products Manufacturing for Measuring, Displaying, and Controlling Industrial Process Variables",
		},

		{
			UserID: 1,
			Name:   "334514 Totalizing Fluid Meter and Counting Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334515 Instrument Manufacturing for Measuring and Testing Electricity and Electrical Signals",
		},

		{
			UserID: 1,
			Name:   "334516 Analytical Laboratory Instrument Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334517 Irradiation Apparatus Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334519 Other Measuring and Controlling Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "334610 Manufacturing and Reproducing Magnetic and Optical Media",
		},

		{
			UserID: 1,
			Name:   "335131 Residential Electric Lighting Fixture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335132 Commercial, Industrial, and Institutional Electric Lighting Fixture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335139 Electric Lamp Bulb and Other Lighting Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335210 Small Electrical Appliance Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335220 Major Household Appliance Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335311 Power, Distribution, and Specialty Transformer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335312 Motor and Generator Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335313 Switchgear and Switchboard Apparatus Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335314 Relay and Industrial Control Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335910 Battery Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335921 Fiber Optic Cable Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335929 Other Communication and Energy Wire Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335931 Current-Carrying Wiring Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335932 Noncurrent-Carrying Wiring Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335991 Carbon and Graphite Product Manufacturing",
		},

		{
			UserID: 1,
			Name:   "335999 All Other Miscellaneous Electrical Equipment and Component Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336110 Automobile and Light Duty Motor Vehicle Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336120 Heavy Duty Truck Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336211 Motor Vehicle Body Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336212 Truck Trailer Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336213 Motor Home Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336214 Travel Trailer and Camper Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336310 Motor Vehicle Gasoline Engine and Engine Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336320 Motor Vehicle Electrical and Electronic Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336330 Motor Vehicle Steering and Suspension Components (except Spring) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336340 Motor Vehicle Brake System Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336350 Motor Vehicle Transmission and Power Train Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336360 Motor Vehicle Seating and Interior Trim Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336370 Motor Vehicle Metal Stamping",
		},

		{
			UserID: 1,
			Name:   "336390 Other Motor Vehicle Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336411 Aircraft Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336412 Aircraft Engine and Engine Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336413 Other Aircraft Parts and Auxiliary Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336414 Guided Missile and Space Vehicle Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336415 Guided Missile and Space Vehicle Propulsion Unit and Propulsion Unit Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336419 Other Guided Missile and Space Vehicle Parts and Auxiliary Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336510 Railroad Rolling Stock Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336611 Ship Building and Repairing",
		},

		{
			UserID: 1,
			Name:   "336612 Boat Building",
		},

		{
			UserID: 1,
			Name:   "336991 Motorcycle, Bicycle, and Parts Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336992 Military Armored Vehicle, Tank, and Tank Component Manufacturing",
		},

		{
			UserID: 1,
			Name:   "336999 All Other Transportation Equipment Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337110 Wood Kitchen Cabinet and Countertop Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337121 Upholstered Household Furniture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337122 Nonupholstered Wood Household Furniture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337126 Household Furniture (except Wood and Upholstered) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337127 Institutional Furniture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337211 Wood Office Furniture Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337212 Custom Architectural Woodwork and Millwork Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337214 Office Furniture (except Wood) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337215 Showcase, Partition, Shelving, and Locker Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337910 Mattress Manufacturing",
		},

		{
			UserID: 1,
			Name:   "337920 Blind and Shade Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339112 Surgical and Medical Instrument Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339113 Surgical Appliance and Supplies Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339114 Dental Equipment and Supplies Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339115 Ophthalmic Goods Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339116 Dental Laboratories",
		},

		{
			UserID: 1,
			Name:   "339910 Jewelry and Silverware Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339920 Sporting and Athletic Goods Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339930 Doll, Toy, and Game Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339940 Office Supplies (except Paper) Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339950 Sign Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339991 Gasket, Packing, and Sealing Device Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339992 Musical Instrument Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339993 Fastener, Button, Needle, and Pin Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339994 Broom, Brush, and Mop Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339995 Burial Casket Manufacturing",
		},

		{
			UserID: 1,
			Name:   "339999 All Other Miscellaneous Manufacturing",
		},

		{
			UserID: 1,
			Name:   "423110 Automobile and Other Motor Vehicle Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423120 Motor Vehicle Supplies and New Parts Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423130 Tire and Tube Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423140 Motor Vehicle Parts (Used) Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423210 Furniture Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423220 Home Furnishing Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423310 Lumber, Plywood, Millwork, and Wood Panel Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423320 Brick, Stone, and Related Construction Material Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423330 Roofing, Siding, and Insulation Material Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423390 Other Construction Material Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423410 Photographic Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423420 Office Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423430 Computer and Computer Peripheral Equipment and Software Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423440 Other Commercial Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423450 Medical, Dental, and Hospital Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423460 Ophthalmic Goods Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423490 Other Professional Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423510 Metal Service Centers and Other Metal Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423520 Coal and Other Mineral and Ore Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423610 Electrical Apparatus and Equipment, Wiring Supplies, and Related Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423620 Household Appliances, Electric Housewares, and Consumer Electronics Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423690 Other Electronic Parts and Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423710 Hardware Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423720 Plumbing and Heating Equipment and Supplies (Hydronics) Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423730 Warm Air Heating and Air-Conditioning Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423740 Refrigeration Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423810 Construction and Mining (except Oil Well) Machinery and Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423820 Farm and Garden Machinery and Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423830 Industrial Machinery and Equipment Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423840 Industrial Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423850 Service Establishment Equipment and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423860 Transportation Equipment and Supplies (except Motor Vehicle) Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423910 Sporting and Recreational Goods and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423920 Toy and Hobby Goods and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423930 Recyclable Material Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423940 Jewelry, Watch, Precious Stone, and Precious Metal Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "423990 Other Miscellaneous Durable Goods Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424110 Printing and Writing Paper Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424120 Stationery and Office Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424130 Industrial and Personal Service Paper Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424210 Drugs and Druggists' Sundries Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424310 Piece Goods, Notions, and Other Dry Goods Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424340 Footwear Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424350 Clothing and Clothing Accessories Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424410 General Line Grocery Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424420 Packaged Frozen Food Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424430 Dairy Product (except Dried or Canned) Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424440 Poultry and Poultry Product Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424450 Confectionery Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424460 Fish and Seafood Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424470 Meat and Meat Product Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424480 Fresh Fruit and Vegetable Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424490 Other Grocery and Related Products Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424510 Grain and Field Bean Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424520 Livestock Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424590 Other Farm Product Raw Material Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424610 Plastics Materials and Basic Forms and Shapes Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424690 Other Chemical and Allied Products Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424710 Petroleum Bulk Stations and Terminals",
		},

		{
			UserID: 1,
			Name:   "424720 Petroleum and Petroleum Products Merchant Wholesalers (except Bulk Stations and Terminals)",
		},

		{
			UserID: 1,
			Name:   "424810 Beer and Ale Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424820 Wine and Distilled Alcoholic Beverage Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424910 Farm Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424920 Book, Periodical, and Newspaper Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424930 Flower, Nursery Stock, and Florists' Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424940 Tobacco Product and Electronic Cigarette Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424950 Paint, Varnish, and Supplies Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "424990 Other Miscellaneous Nondurable Goods Merchant Wholesalers",
		},

		{
			UserID: 1,
			Name:   "425120 Wholesale Trade Agents and Brokers",
		},

		{
			UserID: 1,
			Name:   "441110 New Car Dealers",
		},

		{
			UserID: 1,
			Name:   "441120 Used Car Dealers",
		},

		{
			UserID: 1,
			Name:   "441210 Recreational Vehicle Dealers",
		},

		{
			UserID: 1,
			Name:   "441222 Boat Dealers",
		},

		{
			UserID: 1,
			Name:   "441227 Motorcycle, ATV, and All Other Motor Vehicle Dealers",
		},

		{
			UserID: 1,
			Name:   "441330 Automotive Parts and Accessories Retailers",
		},

		{
			UserID: 1,
			Name:   "441340 Tire Dealers",
		},

		{
			UserID: 1,
			Name:   "444110 Home Centers",
		},

		{
			UserID: 1,
			Name:   "444120 Paint and Wallpaper Retailers",
		},

		{
			UserID: 1,
			Name:   "444140 Hardware Retailers",
		},

		{
			UserID: 1,
			Name:   "444180 Other Building Material Dealers",
		},

		{
			UserID: 1,
			Name:   "444230 Outdoor Power Equipment Retailers",
		},

		{
			UserID: 1,
			Name:   "444240 Nursery, Garden Center, and Farm Supply Retailers",
		},

		{
			UserID: 1,
			Name:   "445110 Supermarkets and Other Grocery Retailers (except Convenience Retailers)",
		},

		{
			UserID: 1,
			Name:   "445131 Convenience Retailers",
		},

		{
			UserID: 1,
			Name:   "445132 Vending Machine Operators",
		},

		{
			UserID: 1,
			Name:   "445230 Fruit and Vegetable Retailers",
		},

		{
			UserID: 1,
			Name:   "445240 Meat Retailers",
		},

		{
			UserID: 1,
			Name:   "445250 Fish and Seafood Retailers",
		},

		{
			UserID: 1,
			Name:   "445291 Baked Goods Retailers",
		},

		{
			UserID: 1,
			Name:   "445292 Confectionery and Nut Retailers",
		},

		{
			UserID: 1,
			Name:   "445298 All Other Specialty Food Retailers",
		},

		{
			UserID: 1,
			Name:   "445320 Beer, Wine, and Liquor Retailers",
		},

		{
			UserID: 1,
			Name:   "449110 Furniture Retailers",
		},

		{
			UserID: 1,
			Name:   "449121 Floor Covering Retailers",
		},

		{
			UserID: 1,
			Name:   "449122 Window Treatment Retailers",
		},

		{
			UserID: 1,
			Name:   "449129 All Other Home Furnishings Retailers",
		},

		{
			UserID: 1,
			Name:   "449210 Electronics and Appliance Retailers",
		},

		{
			UserID: 1,
			Name:   "455110 Department Stores",
		},

		{
			UserID: 1,
			Name:   "455211 Warehouse Clubs and Supercenters",
		},

		{
			UserID: 1,
			Name:   "455219 All Other General Merchandise Retailers",
		},

		{
			UserID: 1,
			Name:   "456110 Pharmacies and Drug Retailers",
		},

		{
			UserID: 1,
			Name:   "456120 Cosmetics, Beauty Supplies, and Perfume Retailers",
		},

		{
			UserID: 1,
			Name:   "456130 Optical Goods Retailers",
		},

		{
			UserID: 1,
			Name:   "456191 Food (Health) Supplement Retailers",
		},

		{
			UserID: 1,
			Name:   "456199 All Other Health and Personal Care Retailers",
		},

		{
			UserID: 1,
			Name:   "457110 Gasoline Stations with Convenience Stores",
		},

		{
			UserID: 1,
			Name:   "457120 Other Gasoline Stations",
		},

		{
			UserID: 1,
			Name:   "457210 Fuel Dealers",
		},

		{
			UserID: 1,
			Name:   "458110 Clothing and Clothing Accessories Retailers",
		},

		{
			UserID: 1,
			Name:   "458210 Shoe Retailers",
		},

		{
			UserID: 1,
			Name:   "458310 Jewelry Retailers",
		},

		{
			UserID: 1,
			Name:   "458320 Luggage and Leather Goods Retailers",
		},

		{
			UserID: 1,
			Name:   "459110 Sporting Goods Retailers",
		},

		{
			UserID: 1,
			Name:   "459120 Hobby, Toy, and Game Retailers",
		},

		{
			UserID: 1,
			Name:   "459130 Sewing, Needlework, and Piece Goods Retailers",
		},

		{
			UserID: 1,
			Name:   "459140 Musical Instrument and Supplies Retailers",
		},

		{
			UserID: 1,
			Name:   "459210 Book Retailers and News Dealers",
		},

		{
			UserID: 1,
			Name:   "459310 Florists",
		},

		{
			UserID: 1,
			Name:   "459410 Office Supplies and Stationery Retailers",
		},

		{
			UserID: 1,
			Name:   "459420 Gift, Novelty, and Souvenir Retailers",
		},

		{
			UserID: 1,
			Name:   "459510 Used Merchandise Retailers",
		},

		{
			UserID: 1,
			Name:   "459910 Pet and Pet Supplies Retailers",
		},

		{
			UserID: 1,
			Name:   "459920 Art Dealers",
		},

		{
			UserID: 1,
			Name:   "459930 Manufactured (Mobile) Home Dealers",
		},

		{
			UserID: 1,
			Name:   "459991 Tobacco, Electronic Cigarette, and Other Smoking Supplies Retailers",
		},

		{
			UserID: 1,
			Name:   "459999 All Other Miscellaneous Retailers",
		},

		{
			UserID: 1,
			Name:   "481111 Scheduled Passenger Air Transportation",
		},

		{
			UserID: 1,
			Name:   "481112 Scheduled Freight Air Transportation",
		},

		{
			UserID: 1,
			Name:   "481211 Nonscheduled Chartered Passenger Air Transportation",
		},

		{
			UserID: 1,
			Name:   "481212 Nonscheduled Chartered Freight Air Transportation",
		},

		{
			UserID: 1,
			Name:   "481219 Other Nonscheduled Air Transportation",
		},

		{
			UserID: 1,
			Name:   "482111 Line-Haul Railroads",
		},

		{
			UserID: 1,
			Name:   "482112 Short Line Railroads",
		},

		{
			UserID: 1,
			Name:   "483111 Deep Sea Freight Transportation",
		},

		{
			UserID: 1,
			Name:   "483112 Deep Sea Passenger Transportation",
		},

		{
			UserID: 1,
			Name:   "483113 Coastal and Great Lakes Freight Transportation",
		},

		{
			UserID: 1,
			Name:   "483114 Coastal and Great Lakes Passenger Transportation",
		},

		{
			UserID: 1,
			Name:   "483211 Inland Water Freight Transportation",
		},

		{
			UserID: 1,
			Name:   "483212 Inland Water Passenger Transportation",
		},

		{
			UserID: 1,
			Name:   "484110 General Freight Trucking, Local",
		},

		{
			UserID: 1,
			Name:   "484121 General Freight Trucking, Long-Distance, Truckload",
		},

		{
			UserID: 1,
			Name:   "484122 General Freight Trucking, Long-Distance, Less Than Truckload",
		},

		{
			UserID: 1,
			Name:   "484210 Used Household and Office Goods Moving",
		},

		{
			UserID: 1,
			Name:   "484220 Specialized Freight (except Used Goods) Trucking, Local",
		},

		{
			UserID: 1,
			Name:   "484230 Specialized Freight (except Used Goods) Trucking, Long-Distance",
		},

		{
			UserID: 1,
			Name:   "485111 Mixed Mode Transit Systems",
		},

		{
			UserID: 1,
			Name:   "485112 Commuter Rail Systems",
		},

		{
			UserID: 1,
			Name:   "485113 Bus and Other Motor Vehicle Transit Systems",
		},

		{
			UserID: 1,
			Name:   "485119 Other Urban Transit Systems",
		},

		{
			UserID: 1,
			Name:   "485210 Interurban and Rural Bus Transportation",
		},

		{
			UserID: 1,
			Name:   "485310 Taxi and Ridesharing Services",
		},

		{
			UserID: 1,
			Name:   "485320 Limousine Service",
		},

		{
			UserID: 1,
			Name:   "485410 School and Employee Bus Transportation",
		},

		{
			UserID: 1,
			Name:   "485510 Charter Bus Industry",
		},

		{
			UserID: 1,
			Name:   "485991 Special Needs Transportation",
		},

		{
			UserID: 1,
			Name:   "485999 All Other Transit and Ground Passenger Transportation",
		},

		{
			UserID: 1,
			Name:   "486110 Pipeline Transportation of Crude Oil",
		},

		{
			UserID: 1,
			Name:   "486210 Pipeline Transportation of Natural Gas",
		},

		{
			UserID: 1,
			Name:   "486910 Pipeline Transportation of Refined Petroleum Products",
		},

		{
			UserID: 1,
			Name:   "486990 All Other Pipeline Transportation",
		},

		{
			UserID: 1,
			Name:   "487110 Scenic and Sightseeing Transportation, Land",
		},

		{
			UserID: 1,
			Name:   "487210 Scenic and Sightseeing Transportation, Water",
		},

		{
			UserID: 1,
			Name:   "487990 Scenic and Sightseeing Transportation, Other",
		},

		{
			UserID: 1,
			Name:   "488111 Air Traffic Control",
		},

		{
			UserID: 1,
			Name:   "488119 Other Airport Operations",
		},

		{
			UserID: 1,
			Name:   "488190 Other Support Activities for Air Transportation",
		},

		{
			UserID: 1,
			Name:   "488210 Support Activities for Rail Transportation",
		},

		{
			UserID: 1,
			Name:   "488310 Port and Harbor Operations",
		},

		{
			UserID: 1,
			Name:   "488320 Marine Cargo Handling",
		},

		{
			UserID: 1,
			Name:   "488330 Navigational Services to Shipping",
		},

		{
			UserID: 1,
			Name:   "488390 Other Support Activities for Water Transportation",
		},

		{
			UserID: 1,
			Name:   "488410 Motor Vehicle Towing",
		},

		{
			UserID: 1,
			Name:   "488490 Other Support Activities for Road Transportation",
		},

		{
			UserID: 1,
			Name:   "488510 Freight Transportation Arrangement",
		},

		{
			UserID: 1,
			Name:   "488991 Packing and Crating",
		},

		{
			UserID: 1,
			Name:   "488999 All Other Support Activities for Transportation",
		},

		{
			UserID: 1,
			Name:   "491110 Postal Service",
		},

		{
			UserID: 1,
			Name:   "492110 Couriers and Express Delivery Services",
		},

		{
			UserID: 1,
			Name:   "492210 Local Messengers and Local Delivery",
		},

		{
			UserID: 1,
			Name:   "493110 General Warehousing and Storage",
		},

		{
			UserID: 1,
			Name:   "493120 Refrigerated Warehousing and Storage",
		},

		{
			UserID: 1,
			Name:   "493130 Farm Product Warehousing and Storage",
		},

		{
			UserID: 1,
			Name:   "493190 Other Warehousing and Storage",
		},

		{
			UserID: 1,
			Name:   "512110 Motion Picture and Video Production",
		},

		{
			UserID: 1,
			Name:   "512120 Motion Picture and Video Distribution",
		},

		{
			UserID: 1,
			Name:   "512131 Motion Picture Theaters (except Drive-Ins)",
		},

		{
			UserID: 1,
			Name:   "512132 Drive-In Motion Picture Theaters",
		},

		{
			UserID: 1,
			Name:   "512191 Teleproduction and Other Postproduction Services",
		},

		{
			UserID: 1,
			Name:   "512199 Other Motion Picture and Video Industries",
		},

		{
			UserID: 1,
			Name:   "512230 Music Publishers",
		},

		{
			UserID: 1,
			Name:   "512240 Sound Recording Studios",
		},

		{
			UserID: 1,
			Name:   "512250 Record Production and Distribution",
		},

		{
			UserID: 1,
			Name:   "512290 Other Sound Recording Industries",
		},

		{
			UserID: 1,
			Name:   "513110 Newspaper Publishers",
		},

		{
			UserID: 1,
			Name:   "513120 Periodical Publishers",
		},

		{
			UserID: 1,
			Name:   "513130 Book Publishers",
		},

		{
			UserID: 1,
			Name:   "513140 Directory and Mailing List Publishers",
		},

		{
			UserID: 1,
			Name:   "513191 Greeting Card Publishers",
		},

		{
			UserID: 1,
			Name:   "513199 All Other Publishers",
		},

		{
			UserID: 1,
			Name:   "513210 Software Publishers",
		},

		{
			UserID: 1,
			Name:   "516110 Radio Broadcasting Stations",
		},

		{
			UserID: 1,
			Name:   "516120 Television Broadcasting Stations",
		},

		{
			UserID: 1,
			Name:   "516210 Media Streaming Distribution Services, Social Networks, and Other Media Networks and Content Providers",
		},

		{
			UserID: 1,
			Name:   "517111 Wired Telecommunications Carriers",
		},

		{
			UserID: 1,
			Name:   "517112 Wireless Telecommunications Carriers (except Satellite)",
		},

		{
			UserID: 1,
			Name:   "517121 Telecommunications Resellers",
		},

		{
			UserID: 1,
			Name:   "517122 Agents for Wireless Telecommunications Services",
		},

		{
			UserID: 1,
			Name:   "517410 Satellite Telecommunications",
		},

		{
			UserID: 1,
			Name:   "517810 All Other Telecommunications",
		},

		{
			UserID: 1,
			Name:   "518210 Computing Infrastructure Providers, Data Processing, Web Hosting, and Related Services",
		},

		{
			UserID: 1,
			Name:   "519210 Libraries and Archives",
		},

		{
			UserID: 1,
			Name:   "519290 Web Search Portals and All Other Information Services",
		},

		{
			UserID: 1,
			Name:   "521110 Monetary Authorities-Central Bank",
		},

		{
			UserID: 1,
			Name:   "522110 Commercial Banking",
		},

		{
			UserID: 1,
			Name:   "522130 Credit Unions",
		},

		{
			UserID: 1,
			Name:   "522180 Savings Institutions and Other Depository Credit Intermediation",
		},

		{
			UserID: 1,
			Name:   "522210 Credit Card Issuing",
		},

		{
			UserID: 1,
			Name:   "522220 Sales Financing",
		},

		{
			UserID: 1,
			Name:   "522291 Consumer Lending",
		},

		{
			UserID: 1,
			Name:   "522292 Real Estate Credit",
		},

		{
			UserID: 1,
			Name:   "522299 International, Secondary Market, and All Other Nondepository Credit Intermediation",
		},

		{
			UserID: 1,
			Name:   "522310 Mortgage and Nonmortgage Loan Brokers",
		},

		{
			UserID: 1,
			Name:   "522320 Financial Transactions Processing, Reserve, and Clearinghouse Activities",
		},

		{
			UserID: 1,
			Name:   "522390 Other Activities Related to Credit Intermediation",
		},

		{
			UserID: 1,
			Name:   "523150 Investment Banking and Securities Intermediation",
		},

		{
			UserID: 1,
			Name:   "523160 Commodity Contracts Intermediation",
		},

		{
			UserID: 1,
			Name:   "523210 Securities and Commodity Exchanges",
		},

		{
			UserID: 1,
			Name:   "523910 Miscellaneous Intermediation",
		},

		{
			UserID: 1,
			Name:   "523940 Portfolio Management and Investment Advice",
		},

		{
			UserID: 1,
			Name:   "523991 Trust, Fiduciary, and Custody Activities",
		},

		{
			UserID: 1,
			Name:   "523999 Miscellaneous Financial Investment Activities",
		},

		{
			UserID: 1,
			Name:   "524113 Direct Life Insurance Carriers",
		},

		{
			UserID: 1,
			Name:   "524114 Direct Health and Medical Insurance Carriers",
		},

		{
			UserID: 1,
			Name:   "524126 Direct Property and Casualty Insurance Carriers",
		},

		{
			UserID: 1,
			Name:   "524127 Direct Title Insurance Carriers",
		},

		{
			UserID: 1,
			Name:   "524128 Other Direct Insurance (except Life, Health, and Medical) Carriers",
		},

		{
			UserID: 1,
			Name:   "524130 Reinsurance Carriers",
		},

		{
			UserID: 1,
			Name:   "524210 Insurance Agencies and Brokerages",
		},

		{
			UserID: 1,
			Name:   "524291 Claims Adjusting",
		},

		{
			UserID: 1,
			Name:   "524292 Pharmacy Benefit Management and Other Third Party Administration of Insurance and Pension Funds",
		},

		{
			UserID: 1,
			Name:   "524298 All Other Insurance Related Activities",
		},

		{
			UserID: 1,
			Name:   "525110 Pension Funds",
		},

		{
			UserID: 1,
			Name:   "525120 Health and Welfare Funds",
		},

		{
			UserID: 1,
			Name:   "525190 Other Insurance Funds",
		},

		{
			UserID: 1,
			Name:   "525910 Open-End Investment Funds",
		},

		{
			UserID: 1,
			Name:   "525920 Trusts, Estates, and Agency Accounts",
		},

		{
			UserID: 1,
			Name:   "525990 Other Financial Vehicles",
		},

		{
			UserID: 1,
			Name:   "531110 Lessors of Residential Buildings and Dwellings",
		},

		{
			UserID: 1,
			Name:   "531120 Lessors of Nonresidential Buildings (except Miniwarehouses)",
		},

		{
			UserID: 1,
			Name:   "531130 Lessors of Miniwarehouses and Self-Storage Units",
		},

		{
			UserID: 1,
			Name:   "531190 Lessors of Other Real Estate Property",
		},

		{
			UserID: 1,
			Name:   "531210 Offices of Real Estate Agents and Brokers",
		},

		{
			UserID: 1,
			Name:   "531311 Residential Property Managers",
		},

		{
			UserID: 1,
			Name:   "531312 Nonresidential Property Managers",
		},

		{
			UserID: 1,
			Name:   "531320 Offices of Real Estate Appraisers",
		},

		{
			UserID: 1,
			Name:   "531390 Other Activities Related to Real Estate",
		},

		{
			UserID: 1,
			Name:   "532111 Passenger Car Rental",
		},

		{
			UserID: 1,
			Name:   "532112 Passenger Car Leasing",
		},

		{
			UserID: 1,
			Name:   "532120 Truck, Utility Trailer, and RV (Recreational Vehicle) Rental and Leasing",
		},

		{
			UserID: 1,
			Name:   "532210 Consumer Electronics and Appliances Rental",
		},

		{
			UserID: 1,
			Name:   "532281 Formal Wear and Costume Rental",
		},

		{
			UserID: 1,
			Name:   "532282 Video Tape and Disc Rental",
		},

		{
			UserID: 1,
			Name:   "532283 Home Health Equipment Rental",
		},

		{
			UserID: 1,
			Name:   "532284 Recreational Goods Rental",
		},

		{
			UserID: 1,
			Name:   "532289 All Other Consumer Goods Rental",
		},

		{
			UserID: 1,
			Name:   "532310 General Rental Centers",
		},

		{
			UserID: 1,
			Name:   "532411 Commercial Air, Rail, and Water Transportation Equipment Rental and Leasing",
		},

		{
			UserID: 1,
			Name:   "532412 Construction, Mining, and Forestry Machinery and Equipment Rental and Leasing",
		},

		{
			UserID: 1,
			Name:   "532420 Office Machinery and Equipment Rental and Leasing",
		},

		{
			UserID: 1,
			Name:   "532490 Other Commercial and Industrial Machinery and Equipment Rental and Leasing",
		},

		{
			UserID: 1,
			Name:   "533110 Lessors of Nonfinancial Intangible Assets (except Copyrighted Works)",
		},

		{
			UserID: 1,
			Name:   "541110 Offices of Lawyers",
		},

		{
			UserID: 1,
			Name:   "541120 Offices of Notaries",
		},

		{
			UserID: 1,
			Name:   "541191 Title Abstract and Settlement Offices",
		},

		{
			UserID: 1,
			Name:   "541199 All Other Legal Services",
		},

		{
			UserID: 1,
			Name:   "541211 Offices of Certified Public Accountants",
		},

		{
			UserID: 1,
			Name:   "541213 Tax Preparation Services",
		},

		{
			UserID: 1,
			Name:   "541214 Payroll Services",
		},

		{
			UserID: 1,
			Name:   "541219 Other Accounting Services",
		},

		{
			UserID: 1,
			Name:   "541310 Architectural Services",
		},

		{
			UserID: 1,
			Name:   "541320 Landscape Architectural Services",
		},

		{
			UserID: 1,
			Name:   "541330 Engineering Services",
		},

		{
			UserID: 1,
			Name:   "541340 Drafting Services",
		},

		{
			UserID: 1,
			Name:   "541350 Building Inspection Services",
		},

		{
			UserID: 1,
			Name:   "541360 Geophysical Surveying and Mapping Services",
		},

		{
			UserID: 1,
			Name:   "541370 Surveying and Mapping (except Geophysical) Services",
		},

		{
			UserID: 1,
			Name:   "541380 Testing Laboratories and Services",
		},

		{
			UserID: 1,
			Name:   "541410 Interior Design Services",
		},

		{
			UserID: 1,
			Name:   "541420 Industrial Design Services",
		},

		{
			UserID: 1,
			Name:   "541430 Graphic Design Services",
		},

		{
			UserID: 1,
			Name:   "541490 Other Specialized Design Services",
		},

		{
			UserID: 1,
			Name:   "541511 Custom Computer Programming Services",
		},

		{
			UserID: 1,
			Name:   "541512 Computer Systems Design Services",
		},

		{
			UserID: 1,
			Name:   "541513 Computer Facilities Management Services",
		},

		{
			UserID: 1,
			Name:   "541519 Other Computer Related Services",
		},

		{
			UserID: 1,
			Name:   "541611 Administrative Management and General Management Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541612 Human Resources Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541613 Marketing Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541614 Process, Physical Distribution, and Logistics Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541618 Other Management Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541620 Environmental Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541690 Other Scientific and Technical Consulting Services",
		},

		{
			UserID: 1,
			Name:   "541713 Research and Development in Nanotechnology",
		},

		{
			UserID: 1,
			Name:   "541714 Research and Development in Biotechnology (except Nanobiotechnology)",
		},

		{
			UserID: 1,
			Name:   "541715 Research and Development in the Physical, Engineering, and Life Sciences (except Nanotechnology and Biotechnology)",
		},

		{
			UserID: 1,
			Name:   "541720 Research and Development in the Social Sciences and Humanities",
		},

		{
			UserID: 1,
			Name:   "541810 Advertising Agencies",
		},

		{
			UserID: 1,
			Name:   "541820 Public Relations Agencies",
		},

		{
			UserID: 1,
			Name:   "541830 Media Buying Agencies",
		},

		{
			UserID: 1,
			Name:   "541840 Media Representatives",
		},

		{
			UserID: 1,
			Name:   "541850 Indoor and Outdoor Display Advertising",
		},

		{
			UserID: 1,
			Name:   "541860 Direct Mail Advertising",
		},

		{
			UserID: 1,
			Name:   "541870 Advertising Material Distribution Services",
		},

		{
			UserID: 1,
			Name:   "541890 Other Services Related to Advertising",
		},

		{
			UserID: 1,
			Name:   "541910 Marketing Research and Public Opinion Polling",
		},

		{
			UserID: 1,
			Name:   "541921 Photography Studios, Portrait",
		},

		{
			UserID: 1,
			Name:   "541922 Commercial Photography",
		},

		{
			UserID: 1,
			Name:   "541930 Translation and Interpretation Services",
		},

		{
			UserID: 1,
			Name:   "541940 Veterinary Services",
		},

		{
			UserID: 1,
			Name:   "541990 All Other Professional, Scientific, and Technical Services",
		},

		{
			UserID: 1,
			Name:   "551111 Offices of Bank Holding Companies",
		},

		{
			UserID: 1,
			Name:   "551112 Offices of Other Holding Companies",
		},

		{
			UserID: 1,
			Name:   "551114 Corporate, Subsidiary, and Regional Managing Offices",
		},

		{
			UserID: 1,
			Name:   "561110 Office Administrative Services",
		},

		{
			UserID: 1,
			Name:   "561210 Facilities Support Services",
		},

		{
			UserID: 1,
			Name:   "561311 Employment Placement Agencies",
		},

		{
			UserID: 1,
			Name:   "561312 Executive Search Services",
		},

		{
			UserID: 1,
			Name:   "561320 Temporary Help Services",
		},

		{
			UserID: 1,
			Name:   "561330 Professional Employer Organizations",
		},

		{
			UserID: 1,
			Name:   "561410 Document Preparation Services",
		},

		{
			UserID: 1,
			Name:   "561421 Telephone Answering Services",
		},

		{
			UserID: 1,
			Name:   "561422 Telemarketing Bureaus and Other Contact Centers",
		},

		{
			UserID: 1,
			Name:   "561431 Private Mail Centers",
		},

		{
			UserID: 1,
			Name:   "561439 Other Business Service Centers (including Copy Shops)",
		},

		{
			UserID: 1,
			Name:   "561440 Collection Agencies",
		},

		{
			UserID: 1,
			Name:   "561450 Credit Bureaus",
		},

		{
			UserID: 1,
			Name:   "561491 Repossession Services",
		},

		{
			UserID: 1,
			Name:   "561492 Court Reporting and Stenotype Services",
		},

		{
			UserID: 1,
			Name:   "561499 All Other Business Support Services",
		},

		{
			UserID: 1,
			Name:   "561510 Travel Agencies",
		},

		{
			UserID: 1,
			Name:   "561520 Tour Operators",
		},

		{
			UserID: 1,
			Name:   "561591 Convention and Visitors Bureaus",
		},

		{
			UserID: 1,
			Name:   "561599 All Other Travel Arrangement and Reservation Services",
		},

		{
			UserID: 1,
			Name:   "561611 Investigation and Personal Background Check Services",
		},

		{
			UserID: 1,
			Name:   "561612 Security Guards and Patrol Services",
		},

		{
			UserID: 1,
			Name:   "561613 Armored Car Services",
		},

		{
			UserID: 1,
			Name:   "561621 Security Systems Services (except Locksmiths)",
		},

		{
			UserID: 1,
			Name:   "561622 Locksmiths",
		},

		{
			UserID: 1,
			Name:   "561710 Exterminating and Pest Control Services",
		},

		{
			UserID: 1,
			Name:   "561720 Janitorial Services",
		},

		{
			UserID: 1,
			Name:   "561730 Landscaping Services",
		},

		{
			UserID: 1,
			Name:   "561740 Carpet and Upholstery Cleaning Services",
		},

		{
			UserID: 1,
			Name:   "561790 Other Services to Buildings and Dwellings",
		},

		{
			UserID: 1,
			Name:   "561910 Packaging and Labeling Services",
		},

		{
			UserID: 1,
			Name:   "561920 Convention and Trade Show Organizers",
		},

		{
			UserID: 1,
			Name:   "561990 All Other Support Services",
		},

		{
			UserID: 1,
			Name:   "562111 Solid Waste Collection",
		},

		{
			UserID: 1,
			Name:   "562112 Hazardous Waste Collection",
		},

		{
			UserID: 1,
			Name:   "562119 Other Waste Collection",
		},

		{
			UserID: 1,
			Name:   "562211 Hazardous Waste Treatment and Disposal",
		},

		{
			UserID: 1,
			Name:   "562212 Solid Waste Landfill",
		},

		{
			UserID: 1,
			Name:   "562213 Solid Waste Combustors and Incinerators",
		},

		{
			UserID: 1,
			Name:   "562219 Other Nonhazardous Waste Treatment and Disposal",
		},

		{
			UserID: 1,
			Name:   "562910 Remediation Services",
		},

		{
			UserID: 1,
			Name:   "562920 Materials Recovery Facilities",
		},

		{
			UserID: 1,
			Name:   "562991 Septic Tank and Related Services",
		},

		{
			UserID: 1,
			Name:   "562998 All Other Miscellaneous Waste Management Services",
		},

		{
			UserID: 1,
			Name:   "611110 Elementary and Secondary Schools",
		},

		{
			UserID: 1,
			Name:   "611210 Junior Colleges",
		},

		{
			UserID: 1,
			Name:   "611310 Colleges, Universities, and Professional Schools",
		},

		{
			UserID: 1,
			Name:   "611410 Business and Secretarial Schools",
		},

		{
			UserID: 1,
			Name:   "611420 Computer Training",
		},

		{
			UserID: 1,
			Name:   "611430 Professional and Management Development Training",
		},

		{
			UserID: 1,
			Name:   "611511 Cosmetology and Barber Schools",
		},

		{
			UserID: 1,
			Name:   "611512 Flight Training",
		},

		{
			UserID: 1,
			Name:   "611513 Apprenticeship Training",
		},

		{
			UserID: 1,
			Name:   "611519 Other Technical and Trade Schools",
		},

		{
			UserID: 1,
			Name:   "611610 Fine Arts Schools",
		},

		{
			UserID: 1,
			Name:   "611620 Sports and Recreation Instruction",
		},

		{
			UserID: 1,
			Name:   "611630 Language Schools",
		},

		{
			UserID: 1,
			Name:   "611691 Exam Preparation and Tutoring",
		},

		{
			UserID: 1,
			Name:   "611692 Automobile Driving Schools",
		},

		{
			UserID: 1,
			Name:   "611699 All Other Miscellaneous Schools and Instruction",
		},

		{
			UserID: 1,
			Name:   "611710 Educational Support Services",
		},

		{
			UserID: 1,
			Name:   "621111 Offices of Physicians (except Mental Health Specialists)",
		},

		{
			UserID: 1,
			Name:   "621112 Offices of Physicians, Mental Health Specialists",
		},

		{
			UserID: 1,
			Name:   "621210 Offices of Dentists",
		},

		{
			UserID: 1,
			Name:   "621310 Offices of Chiropractors",
		},

		{
			UserID: 1,
			Name:   "621320 Offices of Optometrists",
		},

		{
			UserID: 1,
			Name:   "621330 Offices of Mental Health Practitioners (except Physicians)",
		},

		{
			UserID: 1,
			Name:   "621340 Offices of Physical, Occupational and Speech Therapists, and Audiologists",
		},

		{
			UserID: 1,
			Name:   "621391 Offices of Podiatrists",
		},

		{
			UserID: 1,
			Name:   "621399 Offices of All Other Miscellaneous Health Practitioners",
		},

		{
			UserID: 1,
			Name:   "621410 Family Planning Centers",
		},

		{
			UserID: 1,
			Name:   "621420 Outpatient Mental Health and Substance Abuse Centers",
		},

		{
			UserID: 1,
			Name:   "621491 HMO Medical Centers",
		},

		{
			UserID: 1,
			Name:   "621492 Kidney Dialysis Centers",
		},

		{
			UserID: 1,
			Name:   "621493 Freestanding Ambulatory Surgical and Emergency Centers",
		},

		{
			UserID: 1,
			Name:   "621498 All Other Outpatient Care Centers",
		},

		{
			UserID: 1,
			Name:   "621511 Medical Laboratories",
		},

		{
			UserID: 1,
			Name:   "621512 Diagnostic Imaging Centers",
		},

		{
			UserID: 1,
			Name:   "621610 Home Health Care Services",
		},

		{
			UserID: 1,
			Name:   "621910 Ambulance Services",
		},

		{
			UserID: 1,
			Name:   "621991 Blood and Organ Banks",
		},

		{
			UserID: 1,
			Name:   "621999 All Other Miscellaneous Ambulatory Health Care Services",
		},

		{
			UserID: 1,
			Name:   "622110 General Medical and Surgical Hospitals",
		},

		{
			UserID: 1,
			Name:   "622210 Psychiatric and Substance Abuse Hospitals",
		},

		{
			UserID: 1,
			Name:   "622310 Specialty (except Psychiatric and Substance Abuse) Hospitals",
		},

		{
			UserID: 1,
			Name:   "623110 Nursing Care Facilities (Skilled Nursing Facilities)",
		},

		{
			UserID: 1,
			Name:   "623210 Residential Intellectual and Developmental Disability Facilities",
		},

		{
			UserID: 1,
			Name:   "623220 Residential Mental Health and Substance Abuse Facilities",
		},

		{
			UserID: 1,
			Name:   "623311 Continuing Care Retirement Communities",
		},

		{
			UserID: 1,
			Name:   "623312 Assisted Living Facilities for the Elderly",
		},

		{
			UserID: 1,
			Name:   "623990 Other Residential Care Facilities",
		},

		{
			UserID: 1,
			Name:   "624110 Child and Youth Services",
		},

		{
			UserID: 1,
			Name:   "624120 Services for the Elderly and Persons with Disabilities",
		},

		{
			UserID: 1,
			Name:   "624190 Other Individual and Family Services",
		},

		{
			UserID: 1,
			Name:   "624210 Community Food Services",
		},

		{
			UserID: 1,
			Name:   "624221 Temporary Shelters",
		},

		{
			UserID: 1,
			Name:   "624229 Other Community Housing Services",
		},

		{
			UserID: 1,
			Name:   "624230 Emergency and Other Relief Services",
		},

		{
			UserID: 1,
			Name:   "624310 Vocational Rehabilitation Services",
		},

		{
			UserID: 1,
			Name:   "624410 Child Care Services",
		},

		{
			UserID: 1,
			Name:   "711110 Theater Companies and Dinner Theaters",
		},

		{
			UserID: 1,
			Name:   "711120 Dance Companies",
		},

		{
			UserID: 1,
			Name:   "711130 Musical Groups and Artists",
		},

		{
			UserID: 1,
			Name:   "711190 Other Performing Arts Companies",
		},

		{
			UserID: 1,
			Name:   "711211 Sports Teams and Clubs",
		},

		{
			UserID: 1,
			Name:   "711212 Racetracks",
		},

		{
			UserID: 1,
			Name:   "711219 Other Spectator Sports",
		},

		{
			UserID: 1,
			Name:   "711310 Promoters of Performing Arts, Sports, and Similar Events with Facilities",
		},

		{
			UserID: 1,
			Name:   "711320 Promoters of Performing Arts, Sports, and Similar Events without Facilities",
		},

		{
			UserID: 1,
			Name:   "711410 Agents and Managers for Artists, Athletes, Entertainers, and Other Public Figures",
		},

		{
			UserID: 1,
			Name:   "711510 Independent Artists, Writers, and Performers",
		},

		{
			UserID: 1,
			Name:   "712110 Museums",
		},

		{
			UserID: 1,
			Name:   "712120 Historical Sites",
		},

		{
			UserID: 1,
			Name:   "712130 Zoos and Botanical Gardens",
		},

		{
			UserID: 1,
			Name:   "712190 Nature Parks and Other Similar Institutions",
		},

		{
			UserID: 1,
			Name:   "713110 Amusement and Theme Parks",
		},

		{
			UserID: 1,
			Name:   "713120 Amusement Arcades",
		},

		{
			UserID: 1,
			Name:   "713210 Casinos (except Casino Hotels)",
		},

		{
			UserID: 1,
			Name:   "713290 Other Gambling Industries",
		},

		{
			UserID: 1,
			Name:   "713910 Golf Courses and Country Clubs",
		},

		{
			UserID: 1,
			Name:   "713920 Skiing Facilities",
		},

		{
			UserID: 1,
			Name:   "713930 Marinas",
		},

		{
			UserID: 1,
			Name:   "713940 Fitness and Recreational Sports Centers",
		},

		{
			UserID: 1,
			Name:   "713950 Bowling Centers",
		},

		{
			UserID: 1,
			Name:   "713990 All Other Amusement and Recreation Industries",
		},

		{
			UserID: 1,
			Name:   "721110 Hotels (except Casino Hotels) and Motels",
		},

		{
			UserID: 1,
			Name:   "721120 Casino Hotels",
		},

		{
			UserID: 1,
			Name:   "721191 Bed-and-Breakfast Inns",
		},

		{
			UserID: 1,
			Name:   "721199 All Other Traveler Accommodation",
		},

		{
			UserID: 1,
			Name:   "721211 RV (Recreational Vehicle) Parks and Campgrounds",
		},

		{
			UserID: 1,
			Name:   "721214 Recreational and Vacation Camps (except Campgrounds)",
		},

		{
			UserID: 1,
			Name:   "721310 Rooming and Boarding Houses, Dormitories, and Workers' Camps",
		},

		{
			UserID: 1,
			Name:   "722310 Food Service Contractors",
		},

		{
			UserID: 1,
			Name:   "722320 Caterers",
		},

		{
			UserID: 1,
			Name:   "722330 Mobile Food Services",
		},

		{
			UserID: 1,
			Name:   "722410 Drinking Places (Alcoholic Beverages)",
		},

		{
			UserID: 1,
			Name:   "722511 Full-Service Restaurants",
		},

		{
			UserID: 1,
			Name:   "722513 Limited-Service Restaurants",
		},

		{
			UserID: 1,
			Name:   "722514 Cafeterias, Grill Buffets, and Buffets",
		},

		{
			UserID: 1,
			Name:   "722515 Snack and Nonalcoholic Beverage Bars",
		},

		{
			UserID: 1,
			Name:   "811111 General Automotive Repair",
		},

		{
			UserID: 1,
			Name:   "811114 Specialized Automotive Repair",
		},

		{
			UserID: 1,
			Name:   "811121 Automotive Body, Paint, and Interior Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811122 Automotive Glass Replacement Shops",
		},

		{
			UserID: 1,
			Name:   "811191 Automotive Oil Change and Lubrication Shops",
		},

		{
			UserID: 1,
			Name:   "811192 Car Washes",
		},

		{
			UserID: 1,
			Name:   "811198 All Other Automotive Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811210 Electronic and Precision Equipment Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811310 Commercial and Industrial Machinery and Equipment (except Automotive and Electronic) Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811411 Home and Garden Equipment Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811412 Appliance Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "811420 Reupholstery and Furniture Repair",
		},

		{
			UserID: 1,
			Name:   "811430 Footwear and Leather Goods Repair",
		},

		{
			UserID: 1,
			Name:   "811490 Other Personal and Household Goods Repair and Maintenance",
		},

		{
			UserID: 1,
			Name:   "812111 Barber Shops",
		},

		{
			UserID: 1,
			Name:   "812112 Beauty Salons",
		},

		{
			UserID: 1,
			Name:   "812113 Nail Salons",
		},

		{
			UserID: 1,
			Name:   "812191 Diet and Weight Reducing Centers",
		},

		{
			UserID: 1,
			Name:   "812199 Other Personal Care Services",
		},

		{
			UserID: 1,
			Name:   "812210 Funeral Homes and Funeral Services",
		},

		{
			UserID: 1,
			Name:   "812220 Cemeteries and Crematories",
		},

		{
			UserID: 1,
			Name:   "812310 Coin-Operated Laundries and Drycleaners",
		},

		{
			UserID: 1,
			Name:   "812320 Drycleaning and Laundry Services (except Coin-Operated)",
		},

		{
			UserID: 1,
			Name:   "812331 Linen Supply",
		},

		{
			UserID: 1,
			Name:   "812332 Industrial Launderers",
		},

		{
			UserID: 1,
			Name:   "812910 Pet Care (except Veterinary) Services",
		},

		{
			UserID: 1,
			Name:   "812921 Photofinishing Laboratories (except One-Hour)",
		},

		{
			UserID: 1,
			Name:   "812922 One-Hour Photofinishing",
		},

		{
			UserID: 1,
			Name:   "812930 Parking Lots and Garages",
		},

		{
			UserID: 1,
			Name:   "812990 All Other Personal Services",
		},

		{
			UserID: 1,
			Name:   "813110 Religious Organizations",
		},

		{
			UserID: 1,
			Name:   "813211 Grantmaking Foundations",
		},

		{
			UserID: 1,
			Name:   "813212 Voluntary Health Organizations",
		},

		{
			UserID: 1,
			Name:   "813219 Other Grantmaking and Giving Services",
		},

		{
			UserID: 1,
			Name:   "813311 Human Rights Organizations",
		},

		{
			UserID: 1,
			Name:   "813312 Environment, Conservation and Wildlife Organizations",
		},

		{
			UserID: 1,
			Name:   "813319 Other Social Advocacy Organizations",
		},

		{
			UserID: 1,
			Name:   "813410 Civic and Social Organizations",
		},

		{
			UserID: 1,
			Name:   "813910 Business Associations",
		},

		{
			UserID: 1,
			Name:   "813920 Professional Organizations",
		},

		{
			UserID: 1,
			Name:   "813930 Labor Unions and Similar Labor Organizations",
		},

		{
			UserID: 1,
			Name:   "813940 Political Organizations",
		},

		{
			UserID: 1,
			Name:   "813990 Other Similar Organizations (except Business, Professional, Labor, and Political Organizations)",
		},

		{
			UserID: 1,
			Name:   "814110 Private Households",
		},

		{
			UserID: 1,
			Name:   "921110 Executive Offices",
		},

		{
			UserID: 1,
			Name:   "921120 Legislative Bodies",
		},

		{
			UserID: 1,
			Name:   "921130 Public Finance Activities",
		},

		{
			UserID: 1,
			Name:   "921140 Executive and Legislative Offices, Combined",
		},

		{
			UserID: 1,
			Name:   "921150 American Indian and Alaska Native Tribal Governments",
		},

		{
			UserID: 1,
			Name:   "921190 Other General Government Support",
		},

		{
			UserID: 1,
			Name:   "922110 Courts",
		},

		{
			UserID: 1,
			Name:   "922120 Police Protection",
		},

		{
			UserID: 1,
			Name:   "922130 Legal Counsel and Prosecution",
		},

		{
			UserID: 1,
			Name:   "922140 Correctional Institutions",
		},

		{
			UserID: 1,
			Name:   "922150 Parole Offices and Probation Offices",
		},

		{
			UserID: 1,
			Name:   "922160 Fire Protection",
		},

		{
			UserID: 1,
			Name:   "922190 Other Justice, Public Order, and Safety Activities",
		},

		{
			UserID: 1,
			Name:   "923110 Administration of Education Programs",
		},

		{
			UserID: 1,
			Name:   "923120 Administration of Public Health Programs",
		},

		{
			UserID: 1,
			Name:   "923130 Administration of Human Resource Programs (except Education, Public Health, and Veterans' Affairs Programs)",
		},

		{
			UserID: 1,
			Name:   "923140 Administration of Veterans' Affairs",
		},

		{
			UserID: 1,
			Name:   "924110 Administration of Air and Water Resource and Solid Waste Management Programs",
		},

		{
			UserID: 1,
			Name:   "924120 Administration of Conservation Programs",
		},

		{
			UserID: 1,
			Name:   "925110 Administration of Housing Programs",
		},

		{
			UserID: 1,
			Name:   "925120 Administration of Urban Planning and Community and Rural Development",
		},

		{
			UserID: 1,
			Name:   "926110 Administration of General Economic Programs",
		},

		{
			UserID: 1,
			Name:   "926120 Regulation and Administration of Transportation Programs",
		},

		{
			UserID: 1,
			Name:   "926130 Regulation and Administration of Communications, Electric, Gas, and Other Utilities",
		},

		{
			UserID: 1,
			Name:   "926140 Regulation of Agricultural Marketing and Commodities",
		},

		{
			UserID: 1,
			Name:   "926150 Regulation, Licensing, and Inspection of Miscellaneous Commercial Sectors",
		},

		{
			UserID: 1,
			Name:   "927110 Space Research and Technology",
		},

		{
			UserID: 1,
			Name:   "928110 National Security",
		},

		{
			UserID: 1,
			Name:   "928120 International Affairs",
		},
	}

	err := db.Create(&Naicses).Error
	if err != nil {
		log.Fatal(err)
	}
}
