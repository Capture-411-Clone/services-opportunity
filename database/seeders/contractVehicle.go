package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func ContractVehiclesSeeder(db *gorm.DB) {
	vehicle := []models.ContractVehicle{

		{
			UserID: 1,
			Name:   "A&AS (CATS) (AFMSCC - MEDICAL FIELD CONSULTANT ADVISORY AND TECHNICAL SERVICES",
		},

		{
			UserID: 1,
			Name:   "A&AS SETA - ADVISORY AND ASSISTANCE SERVICES FOR SYSTEMS ENGINEERING AND TECHNICAL ASSISTANCE",
		},

		{
			UserID: 1,
			Name:   "A&E II IDIQ - AEC GLOBAL ARCHITECT ENGINEER SERVICES II INDEFINITE DELIVERY INDEFINITE QUANTITY",
		},

		{
			UserID: 1,
			Name:   "A-E13 DCS - AEC ARCHITECT ENGINEERING 2013 DESIGN AND CONSTRUCTION SERVICES",
		},

		{
			UserID: 1,
			Name:   "A2SB - ALLIANT 2 SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "ABMS JADC2 - BROAD AGENCY ANNOUNCEMENT FOR ADVANCED BATTLE MANAGEMENT SYSTEMS",
		},

		{
			UserID: 1,
			Name:   "ACCENT - ARMY CLOUD COMPUTING ENTERPRISE TRANSFORMATION",
		},

		{
			UserID: 1,
			Name:   "ACES - AIRCRAFT MAINTENANCE ENTERPRISE SOLUTION",
		},

		{
			UserID: 1,
			Name:   "ADVISOR - AGILE DELIVERY OF VA IMMINENT STRATEGIC AND OPERATIONAL REQUIREMENTS",
		},

		{
			UserID: 1,
			Name:   "AEC ARCHITECT ENGINEER INDEFINITE DELIVERY INDEFINITE QUANTITY FOR VA NEW ENGLAND HEALTHCARE SYSTEM VISN 1",
		},

		{
			UserID: 1,
			Name:   "AEC ARCHITECT ENGINEER SERVICES FOR THE DESIGN AND OTHER PROFESSIONAL SERVICES NECESSARY TO REHABILITATE MODERNIZE AND DEVELOP NEW OR EXISTING MECHANICAL SYSTEMS FOR FACILITIES AND GROUND SUPPORT",
		},

		{
			UserID: 1,
			Name:   "AEC DEPARTMENT OF HOMELAND SECURITY NATIONAL MULTIPLE AWARD GENERAL ARCHITECT ENGINEERING SERVICES IDIQ CONTRACT",
		},

		{
			UserID: 1,
			Name:   "AEC GENERAL CONSTRUCTION MATOC",
		},

		{
			UserID: 1,
			Name:   "AEC LARGE MACC NAVFAC SOUTHEAST AREA OF RESPONSIBILITY",
		},

		{
			UserID: 1,
			Name:   "AEC NCR MULTIDISCIPLINARY AE DESIGN SERVICES",
		},

		{
			UserID: 1,
			Name:   "AEWD - ADVANCED EXPEDITIONARY WARFARE DEVELOPMENT",
		},

		{
			UserID: 1,
			Name:   "AFCAP IV - AIR FORCE CONTRACT AUGMENTATION PROGRAM IV",
		},

		{
			UserID: 1,
			Name:   "AFICA TRANSIENT ALERT SERVICES MULTIPLE AWARD TASK ORDER CONTRACT",
		},

		{
			UserID: 1,
			Name:   "AFMSA - FACILITY MAINTENANCE SERVICES IN SUPPORT OF AIR FORCE MEDICAL SUPPORT AGENCY",
		},

		{
			UserID: 1,
			Name:   "AIR 6.6 - AIR 6 6 LOGISTICS SUPPORT",
		},

		{
			UserID: 1,
			Name:   "ALLIANT - ALLIANT FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "ALLIANT 2 - ALLIANT 2 UNRESTRICTED",
		},

		{
			UserID: 1,
			Name:   "ALLIANT SB - ALLIANT SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "AMCOM EXPRESS - AMCOM EXPEDITED PROFESSIONAL AND ENGINEERING SUPPORT SERVICES 2004 UMBRELLA",
		},

		{
			UserID: 1,
			Name:   "ASSET FORFEITURE PROGRAM ADMINISTRATIVE SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "ASTRO",
		},

		{
			UserID: 1,
			Name:   "ATSP3 - ADVANCED TECHNOLOGY SUPPORT PROGRAM III",
		},

		{
			UserID: 1,
			Name:   "ATSP4 - ADVANCED TECHNOLOGY SUPPORT PROGRAM IV FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "BAPSLE - BUSINESS AND PROGRAM SOLUTIONS FOR LAW ENFORCEMENT",
		},

		{
			UserID: 1,
			Name:   "C2AD - COMMAND AND CONTROL APPLICATIONS AND INFORMATION SERVICES DEVELOPMENT",
		},

		{
			UserID: 1,
			Name:   "CAAS III - CONTRACTED ADVISORY AND ASSISTANCE SERVICES III",
		},

		{
			UserID: 1,
			Name:   "CAAS IV - CONTRACTED ADVISORY AND ASSISTANCE SERVICES IV",
		},

		{
			UserID: 1,
			Name:   "CDM CMAAS - CONTINUOUS DIAGNOSTICS AND MITIGATION PROGRAM TOOLS AND CONTINUOUS MONITORING AS A SERVICE",
		},

		{
			UserID: 1,
			Name:   "CFT - CONTRACT FIELD TEAM MAINTENANCE",
		},

		{
			UserID: 1,
			Name:   "CHS 4 - COMMON HARDWARE SYSTEMS 4",
		},

		{
			UserID: 1,
			Name:   "CIMS - CENTERS FOR DISEASE CONTROL AND PREVENTION INFORMATION MANAGEMENT SERVICES",
		},

		{
			UserID: 1,
			Name:   "CIO-CS - CHIEF INFORMATION OFFICER COMMODITY SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "CIO-SP2i - CHIEF INFORMATION OFFICER SOLUTIONS AND PARTNER 2 INNOVATIONS",
		},

		{
			UserID: 1,
			Name:   "CIO-SP3 - CHIEF INFORMATION OFFICER SOLUTIONS AND PARTNERS 3 ON RAMP",
		},

		{
			UserID: 1,
			Name:   "CIO-SP3 SB - CHIEF INFORMATION OFFICER SOLUTIONS AND PARTNERS 3 SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "CIO-SP3 UNRESTRICTED - CHIEF INFORMATION OFFICER SOLUTIONS AND PARTNERS 3 UNRESTRICTED",
		},

		{
			UserID: 1,
			Name:   "CLASS)(AFMSCC - CLINICAL ACQUISITION FOR SUPPORT SERVICES PROGRAM",
		},

		{
			UserID: 1,
			Name:   "CLEAN ENERGY FOR NON CRITICAL PRIORITY COUNTRIES",
		},

		{
			UserID: 1,
			Name:   "CME - CYBER MISSION ENGINEERING",
		},

		{
			UserID: 1,
			Name:   "CMMARS - CONTRACTED MAINTENANCE MODIFICATION AIRCREW AND RELATED SERVICES",
		},

		{
			UserID: 1,
			Name:   "CNGT O&L - COUNTER NARCOTICS AND GLOBAL THREATS OPERATIONS AND LOGISTICS",
		},

		{
			UserID: 1,
			Name:   "CNX II - CONNECTIONS II",
		},

		{
			UserID: 1,
			Name:   "COMMIT - CENTRALIZED OPERATIONS AND MAINTENANCE INFORMATION TECHNOLOGY SERVICES",
		},

		{
			UserID: 1,
			Name:   "CRC - CONSUMER RESEARCH AND COMMUNICATION",
		},

		{
			UserID: 1,
			Name:   "CS TATS - CYBER SECURITY AND INFORMATION SYSTEMS TECHNICAL AREA TASKS",
		},

		{
			UserID: 1,
			Name:   "CS2SB - END TO END SOLUTIONS FOR THE FUTURE COMMERCIAL SATCOM ACQUISITION SMALL BUSINESS SET ASIDE",
		},

		{
			UserID: 1,
			Name:   "CS3 - CUSTOM SATCOM SOLUTIONS FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "CTRIC III - COOPERATIVE THREAT REDUCTION PROGRAM SUPPORT CONTRACT III",
		},

		{
			UserID: 1,
			Name:   "CWMD - COMBATING WEAPONS OF MASS DESTRUCTION RESEARCH AND TECHNOLOGY",
		},

		{
			UserID: 1,
			Name:   "DESP III - DESIGN AND ENGINEERING SUPPORT PROGRAM III",
		},

		{
			UserID: 1,
			Name:   "DESS - DISA INFORMATION TECHNOLOGY ENTERPRISE SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "DEVSECOPS BOA - PLATFORM ONE BASIC ORDERING AGREEMENT DEVSECOPS SERVICES",
		},

		{
			UserID: 1,
			Name:   "DHA MEDICAL SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "DISN GSM-O - GLOBAL INFORMATION GRID SERVICES MANAGEMENT OPERATIONS",
		},

		{
			UserID: 1,
			Name:   "DLAE BOA - USACE AEC PETROLEUM FACILITIES MAINTENANCE AND REPAIR SERVICES FOR THE DEFENSE LOGISTICS AGENCY ENERGY PROGRAM BASIC ORDERING AGREEMENT",
		},

		{
			UserID: 1,
			Name:   "DLITE II - DEPARTMENT OF DEFENSE LANGUAGE INTERPRETATION AND TRANSLATION ENTERPRISES II",
		},

		{
			UserID: 1,
			Name:   "DSIDDOMS III - DEFENSE SYSTEMS INTEGRATION DESIGN DEVELOPMENT OPERATION AND MAINTENANCE SUPPORT III",
		},

		{
			UserID: 1,
			Name:   "EAGLE - ENTERPRISE ACQUISITION GATEWAY FOR LEADING EDGE",
		},

		{
			UserID: 1,
			Name:   "EAGLE II - ENHANCED ARMY GLOBAL LOGISTICS ENTERPRISE II",
		},

		{
			UserID: 1,
			Name:   "EAGLE II - ENTERPRISE ACQUISITION GATEWAY FOR LEADING EDGE II FOR SMALL BUSINESSES",
		},

		{
			UserID: 1,
			Name:   "EAGLE II - ENTERPRISE ACQUISITION GATEWAY FOR LEADING EDGE II UNRESTRICTED",
		},

		{
			UserID: 1,
			Name:   "EAGLE) (BOA #1 - ENHANCED ARMY GLOBAL LOGISTICS ENTERPRISE PROGRAM",
		},

		{
			UserID: 1,
			Name:   "EAGLE BOA #7 - ENHANCED ARMY GLOBAL LOGISTICS ENTERPRISE STEP 2",
		},

		{
			UserID: 1,
			Name:   "EIS - ENTERPRISE INFRASTRUCTURE SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "EISM NC2 - NETCENTS II ENTERPRISE INTEGRATION AND SERVICES MANAGEMENT",
		},

		{
			UserID: 1,
			Name:   "ENCORE II - ENCORE II INFORMATION TECHNOLOGY SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "ENCORE III",
		},

		{
			UserID: 1,
			Name:   "ENTERPRISE SOFTWARE SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "ERP - ARMY ENTERPRISE RESOURCE PLANNING SERVICES",
		},

		{
			UserID: 1,
			Name:   "ESPC - ENERGY SAVING PERFORMANCE CONTRACTS",
		},

		{
			UserID: 1,
			Name:   "ESPC - ENERGY SAVINGS PERFORMANCE CONTRACT",
		},

		{
			UserID: 1,
			Name:   "ETS2 - E GOV TRAVEL SERVICES 2",
		},

		{
			UserID: 1,
			Name:   "ETSC - ENTERPRISE TRAINING SERVICES CONTRACT",
		},

		{
			UserID: 1,
			Name:   "EVAL-ME - USAID PPL MONITORING AND EVALUATION",
		},

		{
			UserID: 1,
			Name:   "EVAL-ME - USAID PPL MONITORING EVALUATION AND LEARNING IDIQ",
		},

		{
			UserID: 1,
			Name:   "EWAAC - EGLIN WIDE AGILE ACQUISITION CONTRACT",
		},

		{
			UserID: 1,
			Name:   "F2AST - FUTURE FLEXIBLE ACQUISITION AND SUSTAINMENT TOOL",
		},

		{
			UserID: 1,
			Name:   "FAS - FLEXIBLE ACQUISITION AND SUSTAINMENT FOR THE B 52 WEAPON SYSTEM",
		},

		{
			UserID: 1,
			Name:   "FEDERAL INSURANCE AND MITIGATION ADMINISTRATION NON A AND E SERVICE",
		},

		{
			UserID: 1,
			Name:   "FEDLINK INFORMATION RETRIEVAL SERVICES STRATEGIC SOURCING",
		},

		{
			UserID: 1,
			Name:   "FIA2T - FLEXIBLE INFORMATION ASSURANCE ACQUISITION TOOL",
		},

		{
			UserID: 1,
			Name:   "FIRST - FIELD AND INSTALLATION READINESS SUPPORT TEAM UNRESTRICTED",
		},

		{
			UserID: 1,
			Name:   "FIRSTSOURCE - FIRST SOURCE INFORMATION TECHNOLOGY COMMODITY SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "FIRSTSOURCE II - FIRST SOURCE INFORMATION TECHNOLOGY COMMODITY SOLUTIONS II",
		},

		{
			UserID: 1,
			Name:   "FPS3 - FORCE PROTECTION SECURITY SYSTEMS",
		},

		{
			UserID: 1,
			Name:   "FSSI BMO FO - AEC FEDERAL STRATEGIC SOURCING INITIATIVE BUILDING MAINTENANCE AND OPERATIONS UNRESTRICTED PHASE II ZONES 2 THROUGH 6",
		},

		{
			UserID: 1,
			Name:   "FSSI BMO SB - AEC FEDERAL STRATEGIC SOURCING INITIATIVE BUILDING MAINTENANCE AND OPERATIONS SMALL BUSINESS PHASE II ZONES 2 THROUGH 6",
		},

		{
			UserID: 1,
			Name:   "FSSI BMOS1 - AEC FEDERAL STRATEGIC SOURCING INITIATIVE BUILDING MAINTENANCE AND OPERATIONS PHASE I SMALL BUSINESS ZONE 1 ON RAMP",
		},

		{
			UserID: 1,
			Name:   "FSSI BMOS1 - AEC FEDERAL STRATEGIC SOURCING INITIATIVE BUILDING MAINTENANCE AND OPERATIONS SMALL BUSINESS ZONE 1",
		},

		{
			UserID: 1,
			Name:   "FSSI BMOU1 - AEC FEDERAL STRATEGIC SOURCING INITIATIVE BUILDING MAINTENANCE AND OPERATIONS UNRESTRICTED ZONE 1",
		},

		{
			UserID: 1,
			Name:   "FTSS III - FIELDED TRAINING SYSTEMS SUPPORT III",
		},

		{
			UserID: 1,
			Name:   "FTSS IV - FIELDED TRAINING SYSTEMS SUPPORT IV",
		},

		{
			UserID: 1,
			Name:   "FTSS V - FIELDED TRAINING SYSTEMS SUPPORT V",
		},

		{
			UserID: 1,
			Name:   "GASC COS - GENERAL CONSTRUCTION MATOC SOUTH ATLANTIC DIVISION",
		},

		{
			UserID: 1,
			Name:   "GCSS-AF 2 - GLOBAL COMBAT SUPPORT SYSTEM",
		},

		{
			UserID: 1,
			Name:   "GIG DISN GSM-ETI - GLOBAL INFORMATION GRID SERVICES MANAGEMENT ENGINEERING TRANSITION AND IMPLEMENTATION",
		},

		{
			UserID: 1,
			Name:   "GISS - INSCOM GLOBAL INTELLIGENCE SUPPORT SERVICE IDIQ",
		},

		{
			UserID: 1,
			Name:   "GNS - GLOBAL NETWORK SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA CONSOLIDATED MULTIPLE AWARD SCHEDULE",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 00CORP PSS PROFESSIONAL SERVICES SCHEDULE (AIMS, FABS, ENVIRONMENTAL SERVICES, LANGUAGE SERVICES, LOGWORLD, MOBIS, PES)",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 03FAC FACILITIES MAINTENANCE AND MANAGEMENT",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 36 OFFICE IMAGE AND DOCUMENT SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 48 TRANSPORTATION DELIVERY AND RELOCATION SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 51 V HARDWARE SUPERSTORE",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 520 FABS FINANCIAL AND BUSINESS SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 541 AIMS ADVERTISING AND INTEGRATED MARKETING SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 56 BUILDINGS AND BUILDING MATERIALS INDUSTRIAL SERVICES AND SUPPLIES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 58 I PROFESSIONAL AUDIO VIDEO TELEMETRY TRACKING RECORDING REPRODUCING AND SIGNAL DATA SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 599 TRAVEL SERVICES SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 621 I PROFESSIONAL AND ALLIED HEALTHCARE STAFFING SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 621 II MEDICAL LABORATORY TESTING AND ANALYSIS SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 65 II A MEDICAL EQUIPMENT AND SUPPLIES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 66 SCIENTIFIC EQUIPMENT AND SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 70 IT EQUIPMENT SOFTWARE AND SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 71 II K COMPREHENSIVE FURNITURE MANAGEMENT SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 73 FOOD SERVICE HOSPITALITY CLEANING EQUIPMENT AND SUPPLIES CHEMICALS AND SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 736 TAPS TEMPORARY ADMINISTRATIVE AND PROFESSIONAL STAFFING",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 738X HUMAN RESOURCES AND EQUAL EMPLOYMENT OPPORTUNITY SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 76 PUBLICATION MEDIA",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 78 SPORTS PROMOTIONAL OUTDOOR RECREATION TROPHIES AND SIGNS",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 84 LAW ENFORCEMENT SECURITY MARINE CRAFT FIRE RESCUE AND SPECIAL PURPOSE CLOTHING",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 871 PES PROFESSIONAL ENGINEERING SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 874 MOBIS MISSION ORIENTED BUSINESS INTEGRATED SERVICES",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 874 V LOGWORLD LOGISTICS WORLDWIDE",
		},

		{
			UserID: 1,
			Name:   "GSA SCHEDULE 899 ENVIRONMENTAL ADVISORY SERVICES",
		},

		{
			UserID: 1,
			Name:   "GTACS WTCSS - GLOBAL TACTICAL ADVANCED COMMUNICATION SYSTEMS AND SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "GTSS - GEOSPATIAL TECHNICAL SUPPORT SERVICES DHS WIDE",
		},

		{
			UserID: 1,
			Name:   "HCATS - HUMAN CAPITAL AND TRAINING SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "HCATS SB - HUMAN CAPITAL AND TRAINING SOLUTIONS SMALL BUSINESS SET ASIDE",
		},

		{
			UserID: 1,
			Name:   "HCATS SB 8A - HUMAN CAPITAL AND TRAINING SOLUTIONS 8A SMALL BUSINESS POOL 1 AND POOL 2 ON RAMP",
		},

		{
			UserID: 1,
			Name:   "HICDpro - HUMAN INSTITUTIONAL CAPACITY DEVELOPMENT PARTICIPANT TRAINING",
		},

		{
			UserID: 1,
			Name:   "HITSS - HYBRID INFORMATION TECHNOLOGY SERVICES FOR STATE",
		},

		{
			UserID: 1,
			Name:   "HSOAC FFRDC - HOMELAND SECURITY OPERATIONAL ANALYSIS CENTER FEDERALLY FUNDED RESEARCH AND DEVELOPMENT CENTER",
		},

		{
			UserID: 1,
			Name:   "IAC MAC - DTIC INFORMATION ANALYSIS CENTERS MULTIPLE AWARD CONTRACT",
		},

		{
			UserID: 1,
			Name:   "IDIQ - SAMHSA INDEFINITE DELIVERY INDEFINITE QUANTITY",
		},

		{
			UserID: 1,
			Name:   "ILMS - INSTALLATION AND LOGISTICS MANAGEMENT SERVICES",
		},

		{
			UserID: 1,
			Name:   "IMCS 2 - INFORMATION MANAGEMENT COMMUNICATIONS SERVICES 2",
		},

		{
			UserID: 1,
			Name:   "IMCS III - INFORMATION MANAGEMENT COMMUNICATIONS SERVICES III",
		},

		{
			UserID: 1,
			Name:   "ITES 2S - INFORMATION TECHNOLOGY ENTERPRISE SOLUTIONS 2 SERVICES",
		},

		{
			UserID: 1,
			Name:   "ITES 3H - INFORMATION TECHNOLOGY ENTERPRISE SOLUTIONS 3 HARDWARE",
		},

		{
			UserID: 1,
			Name:   "ITES 3S - INFORMATION TECHNOLOGY ENTERPRISE SOLUTIONS 3 SERVICES",
		},

		{
			UserID: 1,
			Name:   "ITS-SB - INFORMATION TECHNOLOGY SERVICES SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "ITSC - INFORMATION TECHNOLOGY SERVICES CONTRACT",
		},

		{
			UserID: 1,
			Name:   "ITSS - INFORMATION TECHNOLOGIES SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "JETS - DLA J6 ENTERPRISE TECHNOLOGY SERVICES",
		},

		{
			UserID: 1,
			Name:   "LOGCAP IV",
		},

		{
			UserID: 1,
			Name:   "MACC - AEC 8A MULTIPLE AWARD CONSTRUCTION CONTRACT IN NAVFAC WASHINGTON",
		},

		{
			UserID: 1,
			Name:   "MACC II - AEC MULTIPLE AWARD CONSTRUCTION CONTRACT II",
		},

		{
			UserID: 1,
			Name:   "MACC IDIQ - AEC MECHANICAL MULTIPLE AWARD CONSTRUCTION CONTRACT MECHANICAL CONSTRUCTION PROJECTS NORTH CAROLINA AREA",
		},

		{
			UserID: 1,
			Name:   "MACC IDIQ - AEC MULTIPLE AWARD CONSTRUCTION CONTRACT AT HOLLOMAN AFB NEW MEXICO",
		},

		{
			UserID: 1,
			Name:   "MARCORSYSCOM ELECTRONIC AND COMMUNICATION SYSTEMS ECS MATOC",
		},

		{
			UserID: 1,
			Name:   "MATOC - USACE AEC CONSTRUCTION AND INCIDENTAL SERVICE PROJECTS SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "MATOC)(Group IV - AEC REGIONAL INDEFINITE DELIVERY INDEFINITE QUANITITY AND MULTIPLE AWARD TASK ORDER CONTRACTS FOR DREDGING PROJECTS",
		},

		{
			UserID: 1,
			Name:   "MATOC)(IDIQ - AEC GENERAL CONSTRUCTION FOR VISION",
		},

		{
			UserID: 1,
			Name:   "MATOC)(IDIQ)(DPW - AEC CONSTRUCTION SERVICES IN SUPPORT OF THE DIRECTORATE OF PUBLIC WORKS AT FORT BENNING GEORGIA",
		},

		{
			UserID: 1,
			Name:   "MATOC)(IDIQ)(SDVO - AEC CONSTRUCTION SERVICES FOR THE NORTH REGION VISN 8",
		},

		{
			UserID: 1,
			Name:   "MATOC)(IDIQ)(SDVO - AEC CONSTRUCTION SERVICES FOR THE SOUTH REGION VISN 8",
		},

		{
			UserID: 1,
			Name:   "MCOE - INSTRUCTOR AND TRAINING SUPPORT MATOCS FOR THE MANEUVER CENTER OF EXCELLENCE",
		},

		{
			UserID: 1,
			Name:   "MCOE - INSTRUCTOR AND TRAINING SUPPORT SERVICES FOR THE MANEUVER CENTER OF EXCELLENCE",
		},

		{
			UserID: 1,
			Name:   "MILITARY FREEFALL AND STATIC LINE SUPPORT SERVICES FOR NSWC",
		},

		{
			UserID: 1,
			Name:   "MPEC - MISSION PLANNING ENTERPRISE CONTRACT II",
		},

		{
			UserID: 1,
			Name:   "MQS - DHA MEDICAL Q CODED SERVICES",
		},

		{
			UserID: 1,
			Name:   "MRAD - MEDICARE AND MEDICAID RESEARCH AND DEMONSTRATION",
		},

		{
			UserID: 1,
			Name:   "MSS AFMSCC - MEDICAL SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "NAVAIR PM CSS MAC - NAVAIR INITIATIVE PROGRAM MANAGEMENT CONTRACTOR SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "NC2 - NETCENTS II AF NETOPS INFRASTRUCTURE FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "NC2 - NETCENTS II AF NETOPS INFRASTRUCTURE SMALL BUSINESS SET ASIDE",
		},

		{
			UserID: 1,
			Name:   "NC2 - NETCENTS II APPLICATION SERVICES FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "NC2 - NETCENTS II APPLICATION SERVICES SMALL BUSINESS SET ASIDE",
		},

		{
			UserID: 1,
			Name:   "NC2 PRODUCTS - NETCENTS II NETCENTRIC PRODUCTS",
		},

		{
			UserID: 1,
			Name:   "NEPA SUPPORT SERVICES SMALL BUSINESS SET ASIDE",
		},

		{
			UserID: 1,
			Name:   "NETCENTS - AIR FORCE NETWORK CENTRIC SOLUTIONS",
		},

		{
			UserID: 1,
			Name:   "NETWORX ENTERPRISE",
		},

		{
			UserID: 1,
			Name:   "NETWORX UNIVERSAL",
		},

		{
			UserID: 1,
			Name:   "NIHCATS II - NATIONAL INSTITUTES OF HEALTH CONFERENCE AND ADMINISTRATIVE TRAVEL SERVICES",
		},

		{
			UserID: 1,
			Name:   "NLEC-NG - NATIONAL LOCAL EXCHANGE CARRIER SERVICES FOR THE ENTERPRISE NEXT GENERATION",
		},

		{
			UserID: 1,
			Name:   "NMACC - AEC DESIGN BUILD NATIONAL MULTIPLE AWARD CONSTRUCTION CONTRACT",
		},

		{
			UserID: 1,
			Name:   "NMACC II - AEC NATIONAL MULTIPLE AWARD CONSTRUCTION CONTRACT",
		},

		{
			UserID: 1,
			Name:   "NOAALINK - CORE MANAGEMENT COMPONENTS NOAA LINK",
		},

		{
			UserID: 1,
			Name:   "OASIS - ONE ACQUISITION SOLUTION FOR INTEGRATED SERVICES UNRESTRICTED",
		},

		{
			UserID: 1,
			Name:   "OASIS II - OPERATIONAL APPLICATIONS SUPPORT AND INFORMATION SERVICES",
		},

		{
			UserID: 1,
			Name:   "OASIS SB - ONE ACQUISITION SOLUTION FOR INTEGRATED SERVICES SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "OPTARSS II - SERVICES FOR OPERATIONS PLANNING TRAINING AND RESOURCE SUPPORT SERVICES FOR WARFIGHTER OPERATIONS FULL AND OPEN",
		},

		{
			UserID: 1,
			Name:   "OPTARSS II - SERVICES FOR OPERATIONS PLANNING TRAINING AND RESOURCE SUPPORT SERVICES FOR WARFIGHTER OPERATIONS SMALL BUSINESS",
		},

		{
			UserID: 1,
			Name:   "OSP-4 - ORBITAL SERVICES PROGRAM 4",
		},

		{
			UserID: 1,
			Name:   "PA TAC III - PUBLIC ASSISTANCE AND TECHNICAL ASSISTANCE CONTRACTS",
		},

		{
			UserID: 1,
			Name:   "PA-TAC IV - PUBLIC ASSISTANCE TECHNICAL ASSISTANCE CONTRACTS IV",
		},

		{
			UserID: 1,
			Name:   "PACRM - PILOT AND AIRCREW CURRICULUM REVISION AND MAINTENANCE",
		},

		{
			UserID: 1,
			Name:   "PACTS - PROGRAM MANAGEMENT ADMINISTRATIVE CLERICAL AND TECHNICAL SERVICES FOR DHS",
		},

		{
			UserID: 1,
			Name:   "PACTS II - PROGRAM MANAGEMENT ADMINISTRATIVE CLERICAL AND TECHNICAL SERVICES II",
		},

		{
			UserID: 1,
			Name:   "PEITSS - PLATFORM ENGINEERING AND INTEGRATION FOR TACTICAL AND STRATEGIC SYSTEMS",
		},

		{
			UserID: 1,
			Name:   "PICS - PUBLIC INFORMATION AND COMMUNICATIONS SERVICES",
		},

		{
			UserID: 1,
			Name:   "PICS II - PUBLIC INFORMATION AND COMMUNICATIONS SERVICES II",
		},

		{
			UserID: 1,
			Name:   "PMSS - PROGRAM MANAGEMENT SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "PMSS3 - PROGRAM MANAGEMENT SUPPORT SERVICES 3",
		},

		{
			UserID: 1,
			Name:   "PROGRAM SUPPORT CENTER IDIQ TASK ORDER CONTRACTS",
		},

		{
			UserID: 1,
			Name:   "PROGRAM SUPPORT CENTER PSC TASK ORDER CONTRACTS",
		},

		{
			UserID: 1,
			Name:   "PROTECTIVE SECURITY OFFICER SERVICES IN MARYLAND DISTRICT OF COLUMBIA AND VIRGINIA",
		},

		{
			UserID: 1,
			Name:   "PSO - PROTECTIVE SECURITY OFFICER SERVICES FOR MULTIPLE LOCATIONS IN DC MARYLAND AND VIRGINIA",
		},

		{
			UserID: 1,
			Name:   "PSS HRSS - PERSONNEL SUPPORT SERVICES OF HUMAN RESOURCE SERVICES FOR THE UNITED STATES ARMY",
		},

		{
			UserID: 1,
			Name:   "R2-3G - RAPID RESPONSE 3RD GENERATION",
		},

		{
			UserID: 1,
			Name:   "RE&C - REGIONAL ENGINEERING AND CONSTRUCTION",
		},

		{
			UserID: 1,
			Name:   "REMIS - RESEARCH ENGINEERING MISSION INTEGRATION SERVICES",
		},

		{
			UserID: 1,
			Name:   "REPLACE - RESTORING THE ENVIRONMENT THROUGH PROSPERITY LIVELIHOODS AND CONSERVING ECOSYSTEMS",
		},

		{
			UserID: 1,
			Name:   "RES - AEC REMEDIAL ACQUISITION FRAMEWORK REMEDIATION ENVIRONMENTAL SERVICES",
		},

		{
			UserID: 1,
			Name:   "RMACC - AEC REGIONAL REGIONAL MULTIPLE AWARD CONSTRUCTION CONTRACTS AT VARIOUS DHS FACILITIES THROUGHOUT THE US AND ITS TERRITORIES",
		},

		{
			UserID: 1,
			Name:   "RMACC II - REGIONAL MULTIPLE AWARD CONSTRUCTION CONTRACT II",
		},

		{
			UserID: 1,
			Name:   "RMADA - RESEARCH MEASUREMENT ASSESSMENT DESIGN AND ANALYSIS IDIQ",
		},

		{
			UserID: 1,
			Name:   "RMADA 2 - RESEARCH MEASUREMENT ASSESSMENT DESIGN AND ANALYSIS 2",
		},

		{
			UserID: 1,
			Name:   "RS3 - RESPONSIVE STRATEGIC SOURCING FOR SERVICES",
		},

		{
			UserID: 1,
			Name:   "S3 - STRATEGIC SERVICES SOURCING",
		},

		{
			UserID: 1,
			Name:   "SATELLITE - PROFESSIONAL AND TECHNICAL SUPPORT SERVICES CONTRACT VEHICLE SATELLITE DOMAIN",
		},

		{
			UserID: 1,
			Name:   "SBAST - SMALL BUSINESS ACQUISITION AND SUSTAINMENT TOOL",
		},

		{
			UserID: 1,
			Name:   "SBEAS - SMALL BUSINESS ENTERPRISE APPLICATIONS SOLUTIONS IDIQ",
		},

		{
			UserID: 1,
			Name:   "SBS - SUPPORT BASE SERVICES",
		},

		{
			UserID: 1,
			Name:   "SBS - SUPPORT BASE SERVICES 2019",
		},

		{
			UserID: 1,
			Name:   "SCSS - SOF CORE SERVICES SUPPORT FORMERLY SWMS",
		},

		{
			UserID: 1,
			Name:   "SEAPORT NXG - SEAPORT NEXT GENERATION",
		},

		{
			UserID: 1,
			Name:   "SETA III - SYSTEMS ENGINEERING AND TECHNICAL ASSISTANCE",
		},

		{
			UserID: 1,
			Name:   "SETI - DISA SYSTEMS ENGINEERING TECHNOLOGY AND INNOVATION",
		},

		{
			UserID: 1,
			Name:   "SEWP IV - SOLUTIONS FOR ENTERPRISEWIDE PROCUREMENT IV",
		},

		{
			UserID: 1,
			Name:   "SEWP V - SOLUTIONS FOR ENTERPRISEWIDE PROCUREMENT V",
		},

		{
			UserID: 1,
			Name:   "SHIP REPAIR MAC IDIQ NORFOLK VA",
		},

		{
			UserID: 1,
			Name:   "SIISS - SALESFORCE IMPLEMENTATION INTEGRATION AND SUPPORT SERVICES BLANKET PURCHASE AGREEMENT",
		},

		{
			UserID: 1,
			Name:   "SITEC - SITEC SPECIALTY SERVICES",
		},

		{
			UserID: 1,
			Name:   "SOF/PR MACM - SPECIAL OPERATION FORCES MULTIPLE AWARD CONTRACT FOR MODIFICATIONS",
		},

		{
			UserID: 1,
			Name:   "SPARC - STRATEGIC PARTNERS ACQUISITION READINESS CONTRACT",
		},

		{
			UserID: 1,
			Name:   "SPECIALIZED SECURITY TRAINING",
		},

		{
			UserID: 1,
			Name:   "SSES)(NexGen - SOFTWARE AND SYSTEMS ENGINEERING SERVICES NEXT GENERATION",
		},

		{
			UserID: 1,
			Name:   "STARS - STREAMLINED TECHNOLOGY ACQUISITION RESOURCES FOR SERVICES",
		},

		{
			UserID: 1,
			Name:   "STARS II - STREAMLINED TECHNOLOGY ACQUISITION RESOURCES FOR SERVICES II",
		},

		{
			UserID: 1,
			Name:   "STARS III - STREAMLINED TECHNOLOGY ACQUISITION RESOURCES FOR SERVICES III",
		},

		{
			UserID: 1,
			Name:   "STARSS II - SCIENCES TECHNOLOGY AND RESEARCH SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "START III - SUPERFUND TECHNICAL ASSESSMENT AND RESPONSE TEAM CONTRACT SUPPORT",
		},

		{
			UserID: 1,
			Name:   "STOCK 2 - PEO STRI ALL ACQUISITION",
		},

		{
			UserID: 1,
			Name:   "SWIFT 5 - SUPPORT WHICH IMPLEMENTS FAST TRANSITIONS 5",
		},

		{
			UserID: 1,
			Name:   "SeaPort - PROFESSIONAL ADMINISTRATIVE AND MANAGEMENT SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "SeaPort-e - ENGINEERING TECHNICAL AND SUPPORT SERVICES FOR NAVAL WARFARE CENTER AND OTHER NAVSEA FIELD SITES",
		},

		{
			UserID: 1,
			Name:   "T4 - TRANSFORMATION TWENTY ONE TOTAL TECHNOLOGY PROGRAM",
		},

		{
			UserID: 1,
			Name:   "T4NG - TRANSFORMATION TWENTY ONE TOTAL TECHNOLOGY PROGRAM NEXT GENERATION",
		},

		{
			UserID: 1,
			Name:   "TABSS Domain 1 - PROGRAM MANAGEMENT ENGINEERING AND TECHNOLOGY SUPPORT SERVICES DOMAIN ONE",
		},

		{
			UserID: 1,
			Name:   "TABSS Domain 2 - BUSINESS FINANCIAL MANAGEMENT AND AUDIT SUPPORT SERVICES DOMAIN TWO",
		},

		{
			UserID: 1,
			Name:   "TABSS Domain 3 - CONTRACT MANAGEMENT SUPPORT SERVICES DOMAIN THREE",
		},

		{
			UserID: 1,
			Name:   "TACCOM - ENTERPRISE TACTICAL COMMUNICATIONS EQUIPMENT AND SERVICES",
		},

		{
			UserID: 1,
			Name:   "TACCOM II - TACTICAL COMMUNICATIONS EQUIPMENT AND SERVICES II",
		},

		{
			UserID: 1,
			Name:   "TDPC - TRAINING DATA PRODUCTS CONTRACT",
		},

		{
			UserID: 1,
			Name:   "TDSS(E) - TECHNICAL DATA SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "TEAMS - TRICARE EVALUATION ANALYSIS AND MANAGEMENT SUPPORT",
		},

		{
			UserID: 1,
			Name:   "TECHNICAL SERVICES FOR ASPE AHRQ AND ONC",
		},

		{
			UserID: 1,
			Name:   "TEDSS - TECHNICAL ENGINEERING DEVELOPMENT SUPPORT SERVICES AT C3CEN",
		},

		{
			UserID: 1,
			Name:   "TEPS - ENTERPRISE WIDE TECHNICAL ENGINEERING AND PROGRAMMATIC SUPPORT SERVICES",
		},

		{
			UserID: 1,
			Name:   "TIPSS - 4 - TOTAL INFORMATION PROCESSING SUPPORT SERVICES 4",
		},

		{
			UserID: 1,
			Name:   "TLPS - THUNDERBOLT LIFE CYCLE PROGRAM SUPPORT",
		},

		{
			UserID: 1,
			Name:   "TSA III - TRAINING SYSTEMS ACQUISITION III PROGRAM",
		},

		{
			UserID: 1,
			Name:   "TSA IV - TRAINING SYSTEM ACQUISITION IV",
		},

		{
			UserID: 1,
			Name:   "TSC III - TRAINING SYSTEMS CONTRACT III",
		},

		{
			UserID: 1,
			Name:   "TSC IV - TRAINING SYSTEMS CONTRACT IV",
		},

		{
			UserID: 1,
			Name:   "USA CONTACT MULTICHANNEL CONTACT CENTER SERVICES",
		},

		{
			UserID: 1,
			Name:   "USACE AEC HORIZONTAL CONSTRUCTION CONTRACTS IN SUPPORT OF THE FORT WORTH DISTRICT DEPARTMENT OF HOMELAND SECURITY PROGRAMS",
		},

		{
			UserID: 1,
			Name:   "USACE AEC MATOC FOR HORIZONTAL CONSTRUCTION SERVICES IN SUPPORT OF THE SAN DIEGO EL CENTRO YUMA AND TUCSON BORDER PATROL SECTORS AND THE USACE SOUTHWESTERN DIVISION AND SOUTH PACIFIC DIVISION",
		},

		{
			UserID: 1,
			Name:   "USACE AEC NORTHEAST MATOC",
		},

		{
			UserID: 1,
			Name:   "USACE AEC SOUTHEAST UNITED STATES REGIONAL DESIGN BUILD MATOC",
		},

		{
			UserID: 1,
			Name:   "USAFCOEFS - TRAINING STRATEGIES AND DEVELOPMENT FOR US ARMY FIRES CENTER OF EXCELLENCE",
		},

		{
			UserID: 1,
			Name:   "VECTOR - VETERAN ENTERPRISE CONTRACTING FOR TRANSFORMATION AND OPERATIONAL READINESS",
		},

		{
			UserID: 1,
			Name:   "VETS - VETERANS TECHNOLOGY SERVICES",
		},

		{
			UserID: 1,
			Name:   "VETS 2 - VETERANS TECHNOLOGY SERVICES 2",
		},

		{
			UserID: 1,
			Name:   "VICCS - VETERANS INTAKE CONVERSION AND COMMUNICATION SERVICES",
		},

		{
			UserID: 1,
			Name:   "VRM ITSS - VETERANS RELATIONSHIP MANAGEMENT PROGRAM IT SOLUTIONS AND SUPPORT",
		},

		{
			UserID: 1,
			Name:   "WADI - WATER AND DEVELOPMENT IDIQ",
		},

		{
			UserID: 1,
			Name:   "WEATHER PROTECH - PROFESSIONAL SCIENTIFIC AND TECHNICAL SERVICES WEATHER DOMAIN",
		},

		{
			UserID: 1,
			Name:   "WERC09 - AEC WORLDWIDE ENVIRONMENTAL RESTORATION AND CONSTRUCTION",
		},

		{
			UserID: 1,
			Name:   "WF3 - WORKFORCE 30 MANAGED SOLUTION DEFENSE HEALTHCARE MANAGEMENT SYSTEM",
		},

		{
			UserID: 1,
			Name:   "WITS 3 - WASHINGTON INTERAGENCY TELECOMMUNICATIONS SYSTEM 3",
		},

		{
			UserID: 1,
			Name:   "WORLDWIDE PERSONAL PROTECTION SERVICES",
		},

		{
			UserID: 1,
			Name:   "eFAST MOA - ELECTRONIC FAA ACCELERATED AND SIMPLIFIED MASTER ORDERING AGREEMENT",
		},

		{
			UserID: 1,
			Name:   "eFAST MOA - ELECTRONIC FAA ACCELERATED AND SIMPLIFIED MASTER ORDERING AGREEMENT ON RAMP",
		},
	}

	err := db.Create(&vehicle).Error
	if err != nil {
		log.Fatal(err)
	}
}
