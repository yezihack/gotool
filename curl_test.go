package too

import (
	"errors"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPostForm(t *testing.T) {
	Convey("Curl to Post Form", t, func() {
		Convey("url is nil", func() {
			data, err := PostForm("", nil, nil)
			So(err, ShouldBeError, errors.New("url is empty"))
			So(data, ShouldBeNil)
		})
		Convey("Invalid request", func() {
			uri := "https://localhost/reg"
			params := map[string]interface{}{
				"username": "sgfoot",
				"password": "sgfoot",
				"sex":      1,
			}
			headers := map[string]string{
				"token": "62dd45d0-8802-4ad6-a29a-8a4fc0fa63e1",
			}
			data, err := PostForm(uri, params, headers, time.Second*3)
			So(err, ShouldBeError)
			So(data, ShouldBeNil)
		})
		Convey("Valid request", func() {
			uri := "https://tieba.baidu.com/f/commit/thread/add"
			params := map[string]interface{}{
				"content": "bbb",
				"title":   "bbb",
				"tbs":     "279d9d4e8dd774ac1592038175",
			}
			headers := map[string]string{
				"Origin": "https://tieba.baidu.com",
				"Cookie": "BIDUPSID=491A958B488006A1C955EA497D10C605; PSTM=1536535608; TIEBAUID=771c50dd360767f3b2474fe7; TIEBA_USERTYPE=a5de5c0ad21d0c1ab23b4cb9; bdshare_firstime=1573109380497; BAIDUID=1097F8ADD59E5AF271B4FFEABE5EEFB5:FG=1; BDUSS=TUyTEF-eC04S3lLc0wyazRyMGFRWm5DcjdOTXhNSTlMQ3JhOGEzMn4xaXZ3SGxlRVFBQUFBJCQAAAAAAAAAAAEAAABqDY4EeWV6aWhhY2syMDA4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAK8zUl6vM1Jec; MCITY=-131%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; delPer=0; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; BDRCVFR[dG2JNJb_ajR]=mk3SLVN4HKm; BDRCVFR[tox4WRQ4-Km]=mk3SLVN4HKm; PSINO=1; H_PS_PSSID=31909_1462_31672_21100_31660_32045_31779_31714_30824_31844_26350; STOKEN=9ae5e01ddd1d813cb7dcf4cb60018b928df23cc6806c47c8b4b1252971447e44; 76418410_FRSVideoUploadTip=1; st_key_id=17; Hm_lvt_98b9d8c2fd6608d564bf2ac2ae642948=1592038166,1592038175; Hm_lpvt_98b9d8c2fd6608d564bf2ac2ae642948=1592038175; st_data=5ac137828bb9c38b63881e8bca82af6137a7490fd61c4f074eb0808873b4f8c977020ebcf83d144b4f3053c62790048056263d1d5867dbd05e8d8d1567e29812eb75282d672aa29d1cf65f67c3027f8cfda527722dafd3bc1eb5f2aefbb94837cc2d0d0f8b7cb38f0875619e8e04db1aac62ea609cc0f166cb39824bc53c2f54; st_sign=d58a3480; showCardBeforeSign=1",
			}
			data, err := PostForm(uri, params, headers, time.Second*3)
			So(err, ShouldBeNil)
			So(len(data), ShouldEqual, len([]byte("{\"no\":9,\"err_code\":210009,\"error\":null,\"data\":[]}")))
		})
	})
}

func TestPostJson(t *testing.T) {
	Convey("Curl to Post Json", t, func() {
		Convey("url is nil", func() {
			data, err := PostJson("", nil, nil)
			So(err, ShouldBeError, errors.New("url is empty"))
			So(data, ShouldBeNil)
		})
		Convey("Invalid request", func() {
			uri := "https://localhost/reg"
			params := map[string]interface{}{
				"username": "sgfoot",
				"password": "sgfoot",
				"sex":      1,
			}
			headers := map[string]string{
				"token": "62dd45d0-8802-4ad6-a29a-8a4fc0fa63e1",
			}
			data, err := PostJson(uri, params, headers, time.Second*3)
			So(err, ShouldBeError)
			So(data, ShouldBeNil)
		})
		Convey("Valid request", func() {
			uri := "https://tieba.baidu.com/f/commit/thread/add"
			params := map[string]interface{}{
				"content": "bbb",
				"title":   "bbb",
				"tbs":     "279d9d4e8dd774ac1592038175",
			}
			headers := map[string]string{
				"Origin": "https://tieba.baidu.com",
				"Cookie": "BIDUPSID=491A958B488006A1C955EA497D10C605; PSTM=1536535608; TIEBAUID=771c50dd360767f3b2474fe7; TIEBA_USERTYPE=a5de5c0ad21d0c1ab23b4cb9; bdshare_firstime=1573109380497; BAIDUID=1097F8ADD59E5AF271B4FFEABE5EEFB5:FG=1; BDUSS=TUyTEF-eC04S3lLc0wyazRyMGFRWm5DcjdOTXhNSTlMQ3JhOGEzMn4xaXZ3SGxlRVFBQUFBJCQAAAAAAAAAAAEAAABqDY4EeWV6aWhhY2syMDA4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAK8zUl6vM1Jec; MCITY=-131%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; delPer=0; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; BDRCVFR[dG2JNJb_ajR]=mk3SLVN4HKm; BDRCVFR[tox4WRQ4-Km]=mk3SLVN4HKm; PSINO=1; H_PS_PSSID=31909_1462_31672_21100_31660_32045_31779_31714_30824_31844_26350; STOKEN=9ae5e01ddd1d813cb7dcf4cb60018b928df23cc6806c47c8b4b1252971447e44; 76418410_FRSVideoUploadTip=1; st_key_id=17; Hm_lvt_98b9d8c2fd6608d564bf2ac2ae642948=1592038166,1592038175; Hm_lpvt_98b9d8c2fd6608d564bf2ac2ae642948=1592038175; st_data=5ac137828bb9c38b63881e8bca82af6137a7490fd61c4f074eb0808873b4f8c977020ebcf83d144b4f3053c62790048056263d1d5867dbd05e8d8d1567e29812eb75282d672aa29d1cf65f67c3027f8cfda527722dafd3bc1eb5f2aefbb94837cc2d0d0f8b7cb38f0875619e8e04db1aac62ea609cc0f166cb39824bc53c2f54; st_sign=d58a3480; showCardBeforeSign=1",
			}
			data, err := PostJson(uri, params, headers, time.Second*3)
			So(err, ShouldBeNil)
			So(len(data), ShouldEqual, len([]byte("{\"no\":308,\"err_code\":230308,\"error\":\"\\u7528\\u6237\\u6ca1\\u6709\\u6743\\u9650\",\"data\":{\"second_class_id\":\"\"}}")))
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Get Test", t, func() {
		Convey("url is nil", func() {
			data, err := Get("", nil, nil)
			So(err, ShouldBeError, errors.New("url is empty"))
			So(data, ShouldBeNil)
		})
		Convey("Invalid Request", func() {
			data, err := Get("https://localhost", nil, nil)
			So(err, ShouldBeError)
			So(data, ShouldBeNil)
		})
		Convey("Valid Request", func() {
			data, err := Get("https://github.com", nil, nil)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
		})
		Convey("Header, Params And Timeout", func() {
			data, err := Get("https://github.com", map[string]interface{}{
				"username": "yezihack",
			}, map[string]string{
				"Host": "github.com",
			}, time.Second*3)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
		})
	})
}
