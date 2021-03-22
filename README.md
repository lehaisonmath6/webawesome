Tích hợp reactjs vào beego
-1.Trong project reactjs dùng lệnh npm build
-2.Copy tất cả nội dung trong thư mục build vào trong thư mục react (nếu chưa có thì tạo mới)
-3. Trong router.go thêm dòng lệnh 
	web.SetStaticPath("/static", "react/static")
	web.SetStaticPath("/", "react")
-4. Trong controllers/default.go sửa 
    c.TplName = "index.tpl"  => c.TplName = "index.html"

