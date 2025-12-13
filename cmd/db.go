package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/jihanlugas/calendar/app/propertyprice"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/cryption"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func dbUp() {
	log.Println("Running database migrations...")
	dbUpTable()
	dbUpView()
	dbUpListener()
}

func dbUpTable() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Usercompany{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Property{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Propertyprice{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Propertytimeline{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Unit{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Event{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Transaction{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Transactionevent{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.Transactionproduct{})
	if err != nil {
		panic(err)
	}
}

func dbUpView() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Migrator().DropView(model.VIEW_USER)
	if err != nil {
		panic(err)
	}
	vUser := conn.Model(&model.User{}).Unscoped().
		Select("users.*, usercompanies.id as usercompany_id, usercompanies.company_id as company_id, '' as photo_url, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join usercompanies usercompanies on usercompanies.user_id = users.id").
		Joins("left join users u1 on u1.id = users.create_by").
		Joins("left join users u2 on u2.id = users.update_by").
		Where("usercompanies.is_default_company = ? ", true)
	err = conn.Migrator().CreateView(model.VIEW_USER, gorm.ViewOption{
		Replace: true,
		Query:   vUser,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_COMPANY)
	if err != nil {
		panic(err)
	}
	vCompany := conn.Model(&model.Company{}).Unscoped().
		Select("companies.*, '' as photo_url, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join users u1 on u1.id = companies.create_by").
		Joins("left join users u2 on u2.id = companies.update_by")

	err = conn.Migrator().CreateView(model.VIEW_COMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vCompany,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_USERCOMPANY)
	if err != nil {
		panic(err)
	}
	vUsercompany := conn.Model(&model.Usercompany{}).Unscoped().
		Select("usercompanies.*, companies.name as company_name, users.fullname as user_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = usercompanies.company_id").
		Joins("left join users users on users.id = usercompanies.user_id").
		Joins("left join users u1 on u1.id = usercompanies.create_by").
		Joins("left join users u2 on u2.id = usercompanies.update_by")

	err = conn.Migrator().CreateView(model.VIEW_USERCOMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vUsercompany,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_PROPERTY)
	if err != nil {
		panic(err)
	}
	vProperty := conn.Model(&model.Property{}).Unscoped().
		Select("properties.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = properties.company_id").
		Joins("left join users u1 on u1.id = properties.create_by").
		Joins("left join users u2 on u2.id = properties.update_by")

	err = conn.Migrator().CreateView(model.VIEW_PROPERTY, gorm.ViewOption{
		Replace: true,
		Query:   vProperty,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_PROPERTYPRICE)
	if err != nil {
		panic(err)
	}
	vPropertyprice := conn.Model(&model.Propertyprice{}).Unscoped().
		Select("propertyprices.*, companies.name as company_name, properties.name as property_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = propertyprices.company_id").
		Joins("left join properties properties on properties.id = propertyprices.property_id").
		Joins("left join users u1 on u1.id = propertyprices.create_by").
		Joins("left join users u2 on u2.id = propertyprices.update_by")

	err = conn.Migrator().CreateView(model.VIEW_PROPERTYPRICE, gorm.ViewOption{
		Replace: true,
		Query:   vPropertyprice,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_PROPERTYTIMELINE)
	if err != nil {
		panic(err)
	}
	vPropertytimeline := conn.Model(&model.Propertytimeline{}).Unscoped().
		Select("propertytimelines.*, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join users u1 on u1.id = propertytimelines.create_by").
		Joins("left join users u2 on u2.id = propertytimelines.update_by")

	err = conn.Migrator().CreateView(model.VIEW_PROPERTYTIMELINE, gorm.ViewOption{
		Replace: true,
		Query:   vPropertytimeline,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_UNIT)
	if err != nil {
		panic(err)
	}
	vUnit := conn.Model(&model.Unit{}).Unscoped().
		Select("units.*, companies.name as company_name, properties.name as property_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = units.company_id").
		Joins("left join properties properties on properties.id = units.property_id").
		Joins("left join users u1 on u1.id = units.create_by").
		Joins("left join users u2 on u2.id = units.update_by")

	err = conn.Migrator().CreateView(model.VIEW_UNIT, gorm.ViewOption{
		Replace: true,
		Query:   vUnit,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_EVENT)
	if err != nil {
		panic(err)
	}
	vEvent := conn.Model(&model.Event{}).Unscoped().
		Select("events.*, companies.name as company_name, properties.name as property_name, units.name as unit_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = events.company_id").
		Joins("left join properties properties on properties.id = events.property_id").
		Joins("left join units units on units.id = events.unit_id").
		Joins("left join users u1 on u1.id = events.create_by").
		Joins("left join users u2 on u2.id = events.update_by")

	err = conn.Migrator().CreateView(model.VIEW_EVENT, gorm.ViewOption{
		Replace: true,
		Query:   vEvent,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_PRODUCT)
	if err != nil {
		panic(err)
	}
	vProduct := conn.Model(&model.Product{}).Unscoped().
		Select("products.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = products.company_id").
		Joins("left join users u1 on u1.id = products.create_by").
		Joins("left join users u2 on u2.id = products.update_by")

	err = conn.Migrator().CreateView(model.VIEW_PRODUCT, gorm.ViewOption{
		Replace: true,
		Query:   vProduct,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_TRANSACTION)
	if err != nil {
		panic(err)
	}
	vTransaction := conn.Model(&model.Transaction{}).Unscoped().
		Select("transactions.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = transactions.company_id").
		Joins("left join users u1 on u1.id = transactions.create_by").
		Joins("left join users u2 on u2.id = transactions.update_by")

	err = conn.Migrator().CreateView(model.VIEW_TRANSACTION, gorm.ViewOption{
		Replace: true,
		Query:   vTransaction,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_TRANSACTIONEVENT)
	if err != nil {
		panic(err)
	}
	vTransactionevent := conn.Model(&model.Transactionevent{}).Unscoped().
		Select("transactionevents.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = transactionevents.company_id").
		Joins("left join users u1 on u1.id = transactionevents.create_by").
		Joins("left join users u2 on u2.id = transactionevents.update_by")

	err = conn.Migrator().CreateView(model.VIEW_TRANSACTIONEVENT, gorm.ViewOption{
		Replace: true,
		Query:   vTransactionevent,
	})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().DropView(model.VIEW_TRANSACTIONPRODUCT)
	if err != nil {
		panic(err)
	}
	vTransactionproduct := conn.Model(&model.Transactionproduct{}).Unscoped().
		Select("transactionproducts.*, companies.name as company_name, u1.fullname as create_name, u2.fullname as update_name").
		Joins("left join companies companies on companies.id = transactionproducts.company_id").
		Joins("left join users u1 on u1.id = transactionproducts.create_by").
		Joins("left join users u2 on u2.id = transactionproducts.update_by")

	err = conn.Migrator().CreateView(model.VIEW_TRANSACTIONPRODUCT, gorm.ViewOption{
		Replace: true,
		Query:   vTransactionproduct,
	})
	if err != nil {
		panic(err)
	}

}

func dbUpListener() {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	createFunction := `
		CREATE OR REPLACE FUNCTION notify_event_changes()
		RETURNS TRIGGER AS $$
		DECLARE
			payload JSON;
		BEGIN
			IF TG_OP = 'INSERT' THEN
				payload := json_build_object(
					'operation', TG_OP,
					'new', row_to_json(NEW)
				);
		
			ELSIF TG_OP = 'UPDATE' THEN
				payload := json_build_object(
					'operation', TG_OP,
					'old', row_to_json(OLD),
					'new', row_to_json(NEW)
				);
		
			ELSIF TG_OP = 'DELETE' THEN
				payload := json_build_object(
					'operation', TG_OP,
					'old', row_to_json(OLD)
				);
			END IF;
		
			PERFORM pg_notify('event_changes', payload::text);
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
		`
	if err := conn.Exec(createFunction).Error; err != nil {
		panic("failed to create notify_event_changes FUNCTION: " + err.Error())
	}

	dropTrigger := `
		DO $$
		BEGIN
			IF EXISTS (
				SELECT 1 FROM pg_trigger WHERE tgname = 'event_notify_trigger'
			) THEN
				DROP TRIGGER event_notify_trigger ON events;
			END IF;
		END$$;
		`
	if err := conn.Exec(dropTrigger).Error; err != nil {
		panic("failed to drop existing trigger: " + err.Error())
	}

	createTrigger := `
		CREATE TRIGGER event_notify_trigger
		AFTER INSERT OR UPDATE OR DELETE ON events
		FOR EACH ROW EXECUTE FUNCTION notify_event_changes();
		`
	if err := conn.Exec(createTrigger).Error; err != nil {
		panic("failed to create event_notify_trigger: " + err.Error())
	}

	fmt.Println("Listener setup for table events completed.")
}

func dbDown() {
	log.Println("Reverting database migrations...")
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Exec("DROP SCHEMA public CASCADE").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("CREATE SCHEMA public").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO postgres").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO public").Error
	if err != nil {
		panic(err)
	}
}

func dbSeed() {
	log.Println("Seeding the database with initial data start")

	conn, closeConn := db.GetConnection()
	defer closeConn()

	propertypriceRepo := propertyprice.NewRepository()

	tx := conn.Begin()

	adminID := utils.GetUniqueID()
	userID := "f7416f17-884b-46d3-b7db-b90be60a71c5"
	companyID := "fcc18dfc-b0ef-42ef-8036-28503492a2a1"
	property1ID := "db979b45-30e5-4e70-9eec-0cea0089ae12"
	property2ID := "cd94a84e-33bc-43b0-9da0-52a5b6239ed9"

	now := time.Now()

	password, err := cryption.EncryptAES64("123456")
	if err != nil {
		panic(err)
	}

	users := []model.User{
		{
			ID:                adminID,
			Role:              constant.RoleAdmin,
			Email:             "jihanlugas2@gmail.com",
			Username:          "jihanlugas",
			PhoneNumber:       utils.FormatPhoneTo62("6287770333043"),
			Fullname:          "Jihan Lugas",
			Address:           "Jl. Gunung Sahari No. 10, Jakarta Pusat",
			Passwd:            password,
			PassVersion:       1,
			IsActive:          true,
			AccountVerifiedDt: &now,
			CreateBy:          adminID,
			UpdateBy:          adminID,
		},
		{
			ID:                userID,
			Role:              constant.RoleUseradmin,
			Email:             "admindemo@gmail.com",
			Username:          "admindemo",
			PhoneNumber:       utils.FormatPhoneTo62("6287770331234"),
			Fullname:          "Admin Demo",
			Address:           "Jl. Raya Jatinegara No. 10, Jakarta Timur",
			Passwd:            password,
			PassVersion:       1,
			IsActive:          true,
			AccountVerifiedDt: &now,
			CreateBy:          adminID,
			UpdateBy:          adminID,
		},
	}
	tx.Create(&users)

	companies := []model.Company{
		{
			ID:          companyID,
			Name:        "Demo Company",
			Description: "Demo Company Generated",
			Email:       "companydemo@gmail",
			PhoneNumber: utils.FormatPhoneTo62("6287770331234"),
			Address:     "Jl. M.H. Thamrin No. 10, Jakarta Pusat",
			PhotoID:     "",
			CreateBy:    adminID,
			UpdateBy:    adminID,
		},
	}
	tx.Create(&companies)

	usercompanies := []model.Usercompany{
		{
			UserID:           userID,
			CompanyID:        companyID,
			IsDefaultCompany: true,
			IsCreator:        true,
			CreateBy:         adminID,
			UpdateBy:         adminID,
		},
	}
	tx.Create(&usercompanies)

	openTime, _ := time.Parse(constant.FormatTimeLayout, "01:00")  // jam 08 WIB
	closeTime, _ := time.Parse(constant.FormatTimeLayout, "16:00") // jam 23 WIB
	properties := []model.Property{
		{
			ID:          property1ID,
			Name:        "Badminton",
			Description: "Demo Property Generated",
			CompanyID:   companyID,
			OpenTime:    &openTime,
			CloseTime:   &closeTime,
			CreateBy:    adminID,
			UpdateBy:    adminID,
		},
		{
			ID:          property2ID,
			Name:        "Futsal",
			Description: "Demo Property Generated",
			CompanyID:   companyID,
			OpenTime:    &openTime,
			CloseTime:   &closeTime,
			CreateBy:    adminID,
			UpdateBy:    adminID,
		},
	}
	tx.Create(&properties)

	propertytimelines := []model.Propertytimeline{
		{
			ID:                  property1ID,
			DefaultStartDtValue: 6,
			DefaultStartDtUnit:  model.TimeUnitHour,
			DefaultEndDtValue:   6,
			DefaultEndDtUnit:    model.TimeUnitHour,
			MinZoomTimelineHour: 6,
			MaxZoomTimelineHour: 7 * 24, // 7 Hari
			DragSnapMin:         15,     // 15 Minutes
			CreateBy:            adminID,
			UpdateBy:            adminID,
		},
		{
			ID:                  property2ID,
			DefaultStartDtValue: 1,
			DefaultStartDtUnit:  model.TimeUnitDay,
			DefaultEndDtValue:   1,
			DefaultEndDtUnit:    model.TimeUnitDay,
			MinZoomTimelineHour: 6,
			MaxZoomTimelineHour: 7 * 24, // 7 Hari
			DragSnapMin:         15,     // 15 Minutes
			CreateBy:            adminID,
			UpdateBy:            adminID,
		},
	}
	tx.Create(&propertytimelines)

	startTime, _ := time.Parse(constant.FormatTimeLayout, "10:00") // jam 17 WIB
	endTime, _ := time.Parse(constant.FormatTimeLayout, "16:00")   // jam 23 WIB
	propertyprices := []model.Propertyprice{
		{
			CompanyID:  companyID,
			PropertyID: property1ID,
			Priority:   1,
			Weekdays:   pq.Int32Array{0, 1, 2, 3, 4, 5, 6},
			StartTime:  nil,
			EndTime:    nil,
			Price:      10,
			CreateBy:   adminID,
			UpdateBy:   adminID,
		},
		{
			CompanyID:  companyID,
			PropertyID: property1ID,
			Priority:   2,
			Weekdays:   pq.Int32Array{0, 6},
			StartTime:  &startTime,
			EndTime:    &endTime,
			Price:      100,
			CreateBy:   adminID,
			UpdateBy:   adminID,
		},
		{
			CompanyID:  companyID,
			PropertyID: property2ID,
			Priority:   1,
			Weekdays:   pq.Int32Array{0, 1, 2, 3, 4, 5, 6},
			StartTime:  nil,
			EndTime:    nil,
			Price:      1000,
			CreateBy:   adminID,
			UpdateBy:   adminID,
		},
		{
			CompanyID:  companyID,
			PropertyID: property2ID,
			Priority:   2,
			Weekdays:   pq.Int32Array{0, 6},
			StartTime:  &startTime,
			EndTime:    &endTime,
			Price:      10000,
			CreateBy:   adminID,
			UpdateBy:   adminID,
		},
	}
	tx.Create(&propertyprices)

	units := []model.Unit{}
	for i, property := range properties {
		for j := 0; j < (3 + i); j++ {
			unit := model.Unit{
				PropertyID:  property.ID,
				CompanyID:   companyID,
				Name:        fmt.Sprintf("Lapangan %d", j+1),
				Description: fmt.Sprintf("Generated Data Lapangan %d", j+1),
				CreateBy:    adminID,
				UpdateBy:    adminID,
			}
			units = append(units, unit)
		}
	}
	tx.Create(&units)

	events := []model.Event{}
	currProperty := ""
	startDt := time.Now().Truncate(time.Hour * 24)
	for _, unit := range units {
		if currProperty != unit.PropertyID {
			startDt = time.Now().Add(time.Hour * 24 * -3)
		}

		for j := 0; j < 20; j++ {
			endDt := startDt.Add(time.Hour * time.Duration(utils.GetRandomNumber(1, 5)))

			for startDt.Hour() < 8 || endDt.Hour() > 23 || endDt.Hour() < 8 {
				addDuration := time.Hour * time.Duration(utils.GetRandomNumber(1, 3))
				startDt = startDt.Add(addDuration)
				endDt = endDt.Add(addDuration)
			}

			status := constant.EVENT_STATUS_CONFIRM
			rand := utils.GetRandomNumber(1, 30) % 2
			switch rand {
			case 0:
				status = constant.EVENT_STATUS_HOLD
			case 1:
				status = constant.EVENT_STATUS_CONFIRM
			}

			getPriceReq := request.GetPrice{
				PropertyID: unit.PropertyID,
				StartDt:    startDt,
				EndDt:      endDt,
			}
			price, err := propertypriceRepo.GetPrice(tx, getPriceReq)
			if err != nil {
				fmt.Println("ERR => ", err)
			}
			event := model.Event{
				PropertyID:  unit.PropertyID,
				UnitID:      unit.ID,
				CompanyID:   companyID,
				Name:        fmt.Sprintf("Event %d", j+1),
				Description: fmt.Sprintf("Generated Data Event %d", j+1),
				StartDt:     startDt,
				EndDt:       endDt,
				Status:      status,
				Price:       price,
				CreateBy:    adminID,
				UpdateBy:    adminID,
			}

			startDt = endDt.Add(time.Hour * time.Duration(utils.GetRandomNumber(0, 5)))

			events = append(events, event)
		}

	}

	tx.Create(&events)

	products := []model.Product{
		{CompanyID: companyID, Name: "Lee Mineral 600 Ml", Description: "Lee Mineral 600 Ml", Price: 3500, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Lee Mineral 1.500 Ml", Description: "Lee Mineral 1.500 Ml", Price: 8000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Aqua 600 Ml", Description: "Aqua 600 Ml", Price: 3500, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Aqua 1.500 Ml", Description: "Aqua 1.500 Ml", Price: 8000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Pocari Sweet 600 Ml", Description: "Pocari Sweet 600 Ml", Price: 6000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Pocari Sweet 1.500 Ml", Description: "Pocari Sweet 1.500 Ml", Price: 15000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Hydro Coco 600 Ml", Description: "Hydro Coco 600 Ml", Price: 7000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Hydro Coco 1.500 Ml", Description: "Hydro Coco 1.500 Ml", Price: 18000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Kacang Sukro", Description: "Kacang Sukro", Price: 1000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Roti Coklat", Description: "Roti Coklat", Price: 2000, CreateBy: adminID, UpdateBy: adminID},
		{CompanyID: companyID, Name: "Chocolatos", Description: "Chocolatos", Price: 500, CreateBy: adminID, UpdateBy: adminID},
	}

	tx.Create(&products)

	err = tx.Commit().Error
	if err != nil {
		panic(err)
	}

	log.Println("Seeding the database with initial data end")

}

func dbReset() {
	dbDown()
	dbUp()
	dbSeed()
}
