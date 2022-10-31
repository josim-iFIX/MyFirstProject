func InsertBanner(tz *entities.BannerEntity) (int64, bool, error, string) {
    logger.Log.Println("In side Bannermodel")
    db, err := config.ConnectMySqlDb()
    defer db.Close()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return 0, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    for i:=0;i<len(tz.Groupid);i++{
    count,err :=dataAccess.CheckDuplicateBanner(tz)
    if err != nil {
        return 0, false, err, "Something Went Wronggg"
    }
    if count.Total == 0 {
        id, err := dataAccess.InsertBanner(tz)
        if err != nil {
            return 0, false, err, "Something Went Wronggggggggggggggggggggg"
        }
        return id, true, err, ""
    }else{
        return 0, false, nil, "Data Already Exist.Change It."
    }
}
