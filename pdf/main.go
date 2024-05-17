package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

func main() {

	app := fiber.New()
	app.Get("/exportPDF", func(c *fiber.Ctx) error {
		if err := genneratePDF(); err != nil {
			return err
		}
		return c.SendFile("./ไม่มีไม่ให้บวชนะ.pdf")
	})

	app.Listen("127.0.0.1:5000")
}

func genneratePDF() error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("THSarabun", "", "THSarabunNew.ttf")
	pdf.SetFont("THSarabun", "", 16)

	// Title
	pdf.CellFormat(0, 10, "วันที่ ๑ สิงหาคม พ.ศ. ๒๕๖๑", "", 1, "R", false, 0, "")
	pdf.Ln(10)

	// Subject
	pdf.SetFont("THSarabun", "", 14)
	pdf.CellFormat(0, 10, "เรื่อง    ขอประทานการอุปสมบท", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Main Content
	pdf.SetFont("THSarabun", "", 14)
	content := `กราบทูล สมเด็จพระอริยวงศาคตญาณ สมเด็จพระสังฆราช สกลมหาสังฆปริณายก
ด้วยเกล้ากระหม่อม นายชื่อ นามสกุล เลขที่บัตรประจำตัวประชาชน ๑-๒๓๔๕-๖๗๘๙๐-๑ เป็นบุตร นายชื่อ นามสกุล และ นางชื่อ นามสกุล เกิดเมื่อวันอาทิตย์ ขึ้น (หรือ แรม) ๕ ค่ำ เดือน ๒ ปีชวด เวลา ๘.๐๐ น. ตรงกับวันที่ ๕ เดือน มกราคม พ.ศ. ๒๕๓๐ อายุ ๓๘ ปี ส่วนสูง ๑๗๐ เซนติเมตร น้ำหนัก ๗๐ กิโลกรัม ปัจจุบันประกอบอาชีพ ที่ชื่อสถานที่ทำงาน ในตำแหน่งตำแหน่งงาน ประสงค์จะอุปสมบทเป็นพระภิกษุในพระพุทธศาสนา โดยขอประทานพระเมตตาฝ่าพระบาทโปรดทรงเป็นพระอุปัชฌายะ

การนี้ เกล้ากระหม่อมได้ซักซ้อมอุปสมบทวิธีและรายละเอียดเบื้องต้นในการอุปสมบท ทั้งรับการอบรมกับพระมหาคุณณิช คณะไป โดยประสงค์จะอุปสมบทประมาณ ๓๐ วัน จักเป็นวันและเวลาใดสุดแต่จะทรงพระกรุณาโปรด

จึงกราบทูลมาเพื่อทรงทราบและโปรดประทานอุปสมบท

ควรมิควรแล้วแต่จะโปรด

เกล้ากระหม่อม ______________________
(นายชื่อ นามสกุล)`

	pdf.MultiCell(0, 10, content, "", "L", false)
	err := pdf.OutputFileAndClose("ไม่มีไม่ให้บวชนะ.pdf")
	if err != nil {
		return err
	}
	return nil
}
