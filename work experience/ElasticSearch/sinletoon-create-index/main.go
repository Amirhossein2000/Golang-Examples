var IndexName string
var MT sync.Mutex

func GetIndex(client *elastic.Client) (string) {
	t := time.Now().Round(5 * time.Minute).Format("20190703")
	MT.Lock()
	defer MT.Unlock()

	if t != IndexName {
		exist, err := client.IndexExists(IndexName).Do(context.Background())
		if err != nil {
			logger.Error("elasticsearch index exist %v", err.Error())
			return ""
		}
		if !exist {
			_, err := client.CreateIndex(IndexName).BodyString(getMapping()).Do(context.Background())
			if err != nil {
				logger.Error("elasticsearch index create %v", err.Error())
				return ""
			}
			IndexName = t
		}
	}
	return IndexName
}
