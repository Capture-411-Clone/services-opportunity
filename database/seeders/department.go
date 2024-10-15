package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func DepartmentsSeeder(db *gorm.DB) {
	departments := []models.Department{

		{
			UserID:   1,
			MarketID: 1,
			Name:     "AIR FORCE",
		},

		{
			UserID:   1,
			MarketID: 1,
			Name:     "ARMY",
		},

		{
			UserID:   1,
			MarketID: 1,
			Name:     "DEFENSE",
		},

		{
			UserID:   1,
			MarketID: 1,
			Name:     "MARINE CORPS",
		},

		{
			UserID:   1,
			MarketID: 1,
			Name:     "NAVY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ADVISORY COUNCIL ON HISTORIC PRESERVATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "AFRICAN DEVELOPMENT FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "AGENCY FOR INTERNATIONAL DEVELOPMENT",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "AGRICULTURE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "AMERICAN BATTLE MONUMENTS COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "AMERICORPS (CORPORATION FOR NATIONAL AND COMMUNITY SERVICE)",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ARCHITECTURAL AND TRANSPORTATION BARRIERS COMPLIANCE BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ARMED FORCES RETIREMENT HOME",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "BARRY GOLDWATER SCHOLARSHIP AND EXCELLENCE IN EDUCATION FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "CENTRAL INTELLIGENCE AGENCY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "CHEMICAL SAFETY AND HAZARD INVESTIGATION BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COMMERCE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COMMISSION OF FINE ARTS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COMMITTEE FOR PURCHASE FROM PEOPLE WHO ARE BLIND OR SEVERELY DISABLED",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COMMODITY FUTURES TRADING COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "CONGRESS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "CONSUMER FINANCIAL PROTECTION BUREAU",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "CONSUMER PRODUCT SAFETY COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COUNCIL OF THE INSPECTORS GENERAL ON INTEGRITY AND EFFICIENCY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "COURT SERVICES AND OFFENDER SUPERVISION AGENCY FOR THE DISTRICT OF COLUMBIA",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "DEFENSE NUCLEAR FACILITIES SAFETY BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "DELTA REGIONAL AUTHORITY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "DENALI COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "EDUCATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ELECTION ASSISTANCE COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ENERGY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "ENVIRONMENTAL PROTECTION AGENCY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "EQUAL EMPLOYMENT OPPORTUNITY COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "EXECUTIVE OFFICE OF THE PRESIDENT",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "EXPORT-IMPORT BANK OF THE US",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FARM CREDIT ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL COMMUNICATIONS COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL DEPOSIT INSURANCE CORP",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL ELECTION COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL FINANCIAL INSTITUTIONS EXAMINATION COUNCIL",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL HOUSING FINANCE AGENCY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL LABOR RELATIONS AUTHORITY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL MARITIME COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL MEDIATION AND CONCILIATION SERVICE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL MINE SAFETY AND HEALTH REVIEW COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL RESERVE SYSTEM/BOARD OF GOVERNORS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL RETIREMENT THRIFT INVESTMENT BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "FEDERAL TRADE COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "GENERAL SERVICES ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "GULF COAST ECOSYSTEM RESTORATION COUNCIL",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "HARRY S TRUMAN SCHOLARSHIP FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "HEALTH AND HUMAN SERVICES",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "HOMELAND SECURITY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "HOUSING AND URBAN DEVELOPMENT",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "INSTITUTE OF MUSEUM AND LIBRARY SERVICES",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "INTER-AMERICAN FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "INTERIOR",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "INTERNATIONAL BOUNDARY AND WATER COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "JAPAN-US FRIENDSHIP COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "JUSTICE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "LABOR",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "LEGAL SERVICES CORP",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "MARINE MAMMAL COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "MEDICAID AND CHIP PAYMENT AND ACCESS COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "MERIT SYSTEMS PROTECTION BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "MILLENNIUM CHALLENGE CORPORATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "MORRIS K UDALL FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL AERONAUTICS AND SPACE ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL ARCHIVES AND RECORDS ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL CAPITAL PLANNING COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL CREDIT UNION ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL FOUNDATION ON THE ARTS AND THE HUMANITIES",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL GALLERY OF ART",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL LABOR RELATIONS BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL MEDIATION BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL SCIENCE FOUNDATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NATIONAL TRANSPORTATION SAFETY BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "NUCLEAR REGULATORY COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "OCCUPATIONAL SAFETY AND HEALTH REVIEW COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "OFFICE OF GOVERNMENT ETHICS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "OFFICE OF PERSONNEL MANAGEMENT",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "OFFICE OF SPECIAL COUNSEL",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "PEACE CORPS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "PENSION BENEFIT GUARANTY CORP",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "POSTAL REGULATORY COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "PRIVACY AND CIVIL LIBERTIES OVERSIGHT BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "RAILROAD RETIREMENT BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "RESOLUTION FUNDING CORP (REFCORP)",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SECURITIES AND EXCHANGE COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SELECTIVE SERVICE SYSTEM",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SMALL BUSINESS ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SMITHSONIAN INSTITUTION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SOCIAL SECURITY ADMINISTRATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "STATE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "STATE JUSTICE INSTITUTE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "SURFACE TRANSPORTATION BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "TENNESSEE VALLEY AUTHORITY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "TRANSPORTATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "TREASURY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US ABILITYONE COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US AGENCY FOR GLOBAL MEDIA",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US CAPITOL POLICE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US CHEMICAL SAFETY BOARD",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US COURTS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US HOLOCAUST MEMORIAL COUNCIL",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US INSTITUTE OF PEACE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US INTERNATIONAL DEVELOPMENT FINANCE CORPORATION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US INTERNATIONAL TRADE COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US POSTAL SERVICE",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "US TRADE AND DEVELOPMENT AGENCY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "UTAH RECLAMATION MITIGATION AND CONSERVATION COMMISSION",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "VETERANS AFFAIRS",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "WASHINGTON METROPOLITAN AREA TRANSIT AUTHORITY",
		},

		{
			UserID:   1,
			MarketID: 2,
			Name:     "WOODROW WILSON INTERNATIONAL CENTER FOR SCHOLARS",
		},
	}

	err := db.Create(&departments).Error
	if err != nil {
		log.Fatal(err)
	}
}
