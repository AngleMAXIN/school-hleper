package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var html_text = `


<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">


<head id="headerid1">
	<base target='_self'>
	<title>学生个人考试成绩</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">
	<meta http-equiv="keywords" content="湖南强智科技教务系统">
	<meta http-equiv="description" content="湖南强智科技教务系统">
	<meta http-equiv="X-UA-Compatible" content="IE=EmulateIE8" />
<script type="text/javascript" src="/jsxsd/js/jquery-min.js" language="javascript" ></script>
<script type="text/javascript" src="/jsxsd/js/common.js" language="javascript" ></script>
<script type="text/javascript" src="/jsxsd/js/iepngfix_tilebg.js" language="javascript" ></script>
<link href="/jsxsd/framework/images/common.css" rel="stylesheet" type="text/css" />
<link href="/jsxsd/framework/images/blue.css" rel="stylesheet" type="text/css" id="link_theme" />
<link href="/jsxsd/framework/images/workstation.css" rel="stylesheet" type="text/css" />
</head>
<iframe id="notSession" name="notSession" style="display: none;" src=""></iframe>
<script type="text/javascript">
jQuery(document).ready(function(){
	window.setInterval(function(){
		 document.getElementById("notSession").src = "/jsxsd/framework/blankPage.jsp";
	 }, 1000 * 60 * 10);
});
</script>
<body>







<div class="Nsb_pw">
  <div class="Nsb_top">
    <div class="Nsb_top_logo"><table border="0" cellpadding="0" cellspacing="0"><tr><td height="75" valign="middle"><img style="display:;" id="imgLogo" src="/jsxsd/framework/images/index_logo.gif" /></td></tr></table></div>
    <div class="Nsb_top_menu">
        <div id="Top1_divLoginName" class="Nsb_top_menu_nc" style="color: #000000;">马鑫(201601013109)</div>
        <ul>
            <li><a id="TopUserSetting" info="个人设置" href="#" class="Nsb_top_menu_id"><img src="/jsxsd/framework/images/Nsb_top_p1.jpg" id="Top1_imgTopSmallUserPhoto" /></a></li>
            <span></span>


            <li><a id="TopTheme" info="切换主题" href="#" class="Nsb_top_menu_style" onclick="changeDisplay('divTopTheme')"></a></li>
            <span id="Top1_LBackManageSpliter"></span>
            <span id="Top1_LExitSpliter"></span>
            <li id="Top1_LBackManage">
            	<a info="个人设置" href="/jsxsd/grsz/grsz_xggrxx.do?Ves632DSdyV=NEW_JSD_WDZM" class="Nsb_top_menu_options"></a>
            </li>
            <li id="Top1_LExit"><a info="退出" href="javascript:Logout('/jsxsd');"   class="Nsb_top_menu_exit"></a></li>
        </ul>
     </div>
     <div id="divTopTheme" group="topDiv" style="display:none;z-index:10000">
        <div class="Nsb_top_c Nsb_top_c3">
          <ul>
            <li id="theme_blue" theme="blue" index="1" class="Nsb_top_c3_11" onclick="javascript:changeTheme('/jsxsd','blue',this);"></li>
            <li id="theme_green" theme="green" index="2" class="Nsb_top_c3_2"  onclick="javascript:changeTheme('/jsxsd','green',this);"></li>
            <li id="theme_orange" theme="orange" index="3" class="Nsb_top_c3_3" onclick="javascript:changeTheme('/jsxsd','orange',this);"></li>
            <li id="theme_red" theme="red" index="4" class="Nsb_top_c3_4"  onclick="javascript:changeTheme('/jsxsd','red',this);"></li>
            <li id="theme_purple" theme="purple" index="5" class="Nsb_top_c3_5" onclick="javascript:changeTheme('/jsxsd','purple',this);"></li>
            <li id="theme_gray" theme="gray" index="6" class="Nsb_top_c3_6"  onclick="javascript:changeTheme('/jsxsd','gray',this);"></li>
          </ul>
        </div>
     </div>
   </div>
</div>
<div class="Nsb_menu_pw">
  <div class="Nsb_pw">
    <div id="divFirstMenuClass" class="Nsb_menu menu_cn">
      <ul>
         <li class="Nsb_menu_li_now" title="首页">
         	<a class="Nsb_menu_li_h" id="homepage" href="/jsxsd/framework/main.jsp"><span></span></a>
         </li>

         	<li title="我的桌面">
         		<a id="calender_user_schedule" href="/jsxsd/jxzl/jxzl_query?Ves632DSdyV=NEW_XSD_WDZM">我的桌面</a>
         	</li>

         	<li title="学籍成绩">
         		<a id="calender_user_schedule" href="/jsxsd/xsxj/xjxxgl.do?Ves632DSdyV=NEW_XSD_XJCJ">学籍成绩</a>
         	</li>

         	<li title="培养管理">
         		<a id="calender_user_schedule" href="/jsxsd/pyfa/pyfadg_query?Ves632DSdyV=NEW_XSD_PYGL">培养管理</a>
         	</li>

         	<li title="考试报名">
         		<a id="calender_user_schedule" href="/jsxsd/xsks/xsksap_query?Ves632DSdyV=NEW_XSD_KSBM">考试报名</a>
         	</li>

         	<li title="实践环节">
         		<a id="calender_user_schedule" href="/jsxsd/view/syjx/syyy_find.jsp?Ves632DSdyV=NEW_XSD_SJHJ">实践环节</a>
         	</li>

         	<li title="教学评价">
         		<a id="calender_user_schedule" href="/jsxsd/xspj/xspj_find.do?Ves632DSdyV=NEW_XSD_JXPJ">教学评价</a>
         	</li>

      </ul>
    </div>
  </div>
</div>
<div class="Nsb_pw">
	<br/>
	<div class="Nsb_r_title"><a href="/jsxsd/framework/main.jsp"">首页</a> » 考试成绩 » 课程成绩查询 » 查询列表</div>
	<input type="button" id="btn_back" class="button" value="返 回" onclick="window.location.href='/jsxsd/kscj/cjcx_query';"/>

	<br />

	查询条件：全部
	<br />


				<span>已修读<span>175.5</span>学分，</span>
	<table id="dataList" width="100%" border="0" cellspacing="0" cellpadding="0" class="Nsb_r_list Nsb_table">
		<tr>
			<th class="Nsb_r_list_thb" style="width: 35px;">序号</th>
			<th class="Nsb_r_list_thb" style="width: 140px;">开课学期</th>
			<th class="Nsb_r_list_thb" style="width: 110px;">课程编号</th>
			<th class="Nsb_r_list_thb">课程名称</th>
			<th class="Nsb_r_list_thb" style="width: 60px;">成绩</th>
			<th class="Nsb_r_list_thb" style="width: 50px;">学分</th>
			<th class="Nsb_r_list_thb" style="width: 50px;">总学时</th>
			<th class="Nsb_r_list_thb" style="width: 60px;">考核方式</th>
			<th class="Nsb_r_list_thb" style="width: 60px;">课程属性</th>
			<th class="Nsb_r_list_thb" style="width: 120px;">课程性质</th>
		</tr>

		<tr>
			<td>1</td>
			<td>2016-2017-1</td>
			<td align="left">01100110</td>
			<td align="left">地球科学概论</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002258&zcj=72',700,500)">72</a></td>


			<td>4</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>2</td>
			<td>2016-2017-1</td>
			<td align="left">06100111</td>
			<td align="left">高等数学[2-1]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002260&zcj=80',700,500)">80</a></td>


			<td>6</td>
			<td>96</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>3</td>
			<td>2016-2017-1</td>
			<td align="left">06100310</td>
			<td align="left">线性代数</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181003341&zcj=77',700,500)">77</a></td>


			<td>2</td>
			<td>32</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>4</td>
			<td>2016-2017-1</td>
			<td align="left">06220110</td>
			<td align="left">计算机导论</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181003187&zcj=83',700,500)">83</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>5</td>
			<td>2016-2017-1</td>
			<td align="left">06300111</td>
			<td align="left">基础外语[4-1]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002793&zcj=75',700,500)">75</a></td>


			<td>3</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>6</td>
			<td>2016-2017-1</td>
			<td align="left">06400110</td>
			<td align="left">思想道德修养与法律基础</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182004456&zcj=75',700,500)">75</a></td>


			<td>3</td>
			<td>48</td>
			<td>考查</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>7</td>
			<td>2016-2017-1</td>
			<td align="left">06400110</td>
			<td align="left">思想道德修养与法律基础</td>

			<td style=" color:red;"><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181003175&zcj=0',700,500)">0</a></td>


			<td>3</td>
			<td>48</td>
			<td>考查</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>8</td>
			<td>2016-2017-1</td>
			<td align="left">06400210</td>
			<td align="left">毛泽东思想和中国特色社会主义理论体系概论</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002259&zcj=87',700,500)">87</a></td>


			<td>5</td>
			<td>80</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>9</td>
			<td>2016-2017-1</td>
			<td align="left">06500111</td>
			<td align="left">体育[4-1]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171003014&zcj=79',700,500)">79</a></td>


			<td>1</td>
			<td>36</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>10</td>
			<td>2016-2017-1</td>
			<td align="left">20100110</td>
			<td align="left">军训</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002597&zcj=通过',700,500)">通过</a></td>


			<td>2</td>
			<td>2</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>11</td>
			<td>2016-2017-1</td>
			<td align="left">20100210</td>
			<td align="left">军事理论</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620171002256&zcj=92',700,500)">92</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>12</td>
			<td>2016-2017-2</td>
			<td align="left">06100112</td>
			<td align="left">高等数学[2-2]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000600&zcj=84',700,500)">84</a></td>


			<td>5</td>
			<td>80</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>13</td>
			<td>2016-2017-2</td>
			<td align="left">06105011</td>
			<td align="left">大学物理[2-1]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000595&zcj=85',700,500)">85</a></td>


			<td>3</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>14</td>
			<td>2016-2017-2</td>
			<td align="left">06105111</td>
			<td align="left">大学物理实验[2-1]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000593&zcj=75',700,500)">75</a></td>


			<td>0.5</td>
			<td>14</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>15</td>
			<td>2016-2017-2</td>
			<td align="left">06120810</td>
			<td align="left">离散数学</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000597&zcj=80',700,500)">80</a></td>


			<td>4</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>16</td>
			<td>2016-2017-2</td>
			<td align="left">06200110</td>
			<td align="left">C语言程序设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000598&zcj=96',700,500)">96</a></td>


			<td>4</td>
			<td>72</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>17</td>
			<td>2016-2017-2</td>
			<td align="left">06291110</td>
			<td align="left">C语言课程设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172001144&zcj=92',700,500)">92</a></td>


			<td>2</td>
			<td>2</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>18</td>
			<td>2016-2017-2</td>
			<td align="left">06300112</td>
			<td align="left">基础外语[4-2]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172000596&zcj=79',700,500)">79</a></td>


			<td>3</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>19</td>
			<td>2016-2017-2</td>
			<td align="left">06312510</td>
			<td align="left">实用德语</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172001599&zcj=81',700,500)">81</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>公选</td>
			<td>通识教育选修课</td>
		</tr>

		<tr>
			<td>20</td>
			<td>2016-2017-2</td>
			<td align="left">06500112</td>
			<td align="left">体育[4-2]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201620172001489&zcj=79',700,500)">79</a></td>


			<td>1</td>
			<td>36</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>21</td>
			<td>2017-2018-1</td>
			<td align="left">03200110</td>
			<td align="left">电工电子学</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181000879&zcj=80',700,500)">80</a></td>


			<td>4.5</td>
			<td>80</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>22</td>
			<td>2017-2018-1</td>
			<td align="left">06100510</td>
			<td align="left">概率论与数理统计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001434&zcj=65',700,500)">65</a></td>


			<td>3.5</td>
			<td>56</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>23</td>
			<td>2017-2018-1</td>
			<td align="left">06105012</td>
			<td align="left">大学物理[2-2]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001400&zcj=77',700,500)">77</a></td>


			<td>3</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>24</td>
			<td>2017-2018-1</td>
			<td align="left">06105112</td>
			<td align="left">大学物理实验[2-2]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001625&zcj=83',700,500)">83</a></td>


			<td>0.5</td>
			<td>14</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>25</td>
			<td>2017-2018-1</td>
			<td align="left">06221210</td>
			<td align="left">数据结构</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001403&zcj=90',700,500)">90</a></td>


			<td>3.5</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>26</td>
			<td>2017-2018-1</td>
			<td align="left">06264310</td>
			<td align="left">C++程序设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001402&zcj=92',700,500)">92</a></td>


			<td>2.5</td>
			<td>56</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>27</td>
			<td>2017-2018-1</td>
			<td align="left">06300113</td>
			<td align="left">基础外语[4-3]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001399&zcj=71',700,500)">71</a></td>


			<td>3</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>28</td>
			<td>2017-2018-1</td>
			<td align="left">06310210</td>
			<td align="left">大学实用英语语法</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181002507&zcj=60',700,500)">60</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>公选</td>
			<td>通识教育选修课</td>
		</tr>

		<tr>
			<td>29</td>
			<td>2017-2018-1</td>
			<td align="left">06400310</td>
			<td align="left">马克思主义基本原理</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181001428&zcj=76',700,500)">76</a></td>


			<td>3</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>30</td>
			<td>2017-2018-1</td>
			<td align="left">06500113</td>
			<td align="left">体育[4-3]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181002308&zcj=82',700,500)">82</a></td>


			<td>1</td>
			<td>36</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>31</td>
			<td>2017-2018-1</td>
			<td align="left">20200410</td>
			<td align="left">创新中国</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720181002410&zcj=94',700,500)">94</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>公选</td>
			<td>通识教育选修课</td>
		</tr>

		<tr>
			<td>32</td>
			<td>2017-2018-2</td>
			<td align="left">06215710</td>
			<td align="left">网站建设</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000486&zcj=84',700,500)">84</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>33</td>
			<td>2017-2018-2</td>
			<td align="left">06220210</td>
			<td align="left">Java语言程序设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000485&zcj=84',700,500)">84</a></td>


			<td>4</td>
			<td>80</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>34</td>
			<td>2017-2018-2</td>
			<td align="left">06220310</td>
			<td align="left">计算机组成原理</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000484&zcj=78',700,500)">78</a></td>


			<td>4</td>
			<td>72</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>35</td>
			<td>2017-2018-2</td>
			<td align="left">06220510</td>
			<td align="left">计算机网络原理</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000661&zcj=83',700,500)">83</a></td>


			<td>3</td>
			<td>56</td>
			<td>考试</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>36</td>
			<td>2017-2018-2</td>
			<td align="left">06240110</td>
			<td align="left">软件工程</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000480&zcj=93',700,500)">93</a></td>


			<td>2</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>37</td>
			<td>2017-2018-2</td>
			<td align="left">06291210</td>
			<td align="left">Java语言课程设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182001159&zcj=78',700,500)">78</a></td>


			<td>1</td>
			<td>1</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>38</td>
			<td>2017-2018-2</td>
			<td align="left">06291510</td>
			<td align="left">软件工程课程设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182001160&zcj=96',700,500)">96</a></td>


			<td>2</td>
			<td>2</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>39</td>
			<td>2017-2018-2</td>
			<td align="left">06300114</td>
			<td align="left">基础外语[4-4]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182002457&zcj=66',700,500)">66</a></td>


			<td>3</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>40</td>
			<td>2017-2018-2</td>
			<td align="left">06400410</td>
			<td align="left">中国近现代史纲要</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182000660&zcj=79',700,500)">79</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>41</td>
			<td>2017-2018-2</td>
			<td align="left">06500114</td>
			<td align="left">体育[4-4]</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182002254&zcj=79',700,500)">79</a></td>


			<td>1</td>
			<td>36</td>
			<td>考试</td>
			<td>必修</td>
			<td>通识教育课程</td>
		</tr>

		<tr>
			<td>42</td>
			<td>2017-2018-2</td>
			<td align="left">20200210</td>
			<td align="left">大学生创业基础2.0</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201720182002587&zcj=96',700,500)">96</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>公选</td>
			<td>通识教育选修课</td>
		</tr>

		<tr>
			<td>43</td>
			<td>2018-2019-1</td>
			<td align="left">04106410</td>
			<td align="left">用相声演绎中国文化</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191002455&zcj=89',700,500)">89</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>公选</td>
			<td>通识教育选修课</td>
		</tr>

		<tr>
			<td>44</td>
			<td>2018-2019-1</td>
			<td align="left">06130110</td>
			<td align="left">数值分析</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000375&zcj=74',700,500)">74</a></td>


			<td>3.5</td>
			<td>64</td>
			<td>考查</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>45</td>
			<td>2018-2019-1</td>
			<td align="left">06220410</td>
			<td align="left">操作系统</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000373&zcj=84',700,500)">84</a></td>


			<td>3.5</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>学科基础课程</td>
		</tr>

		<tr>
			<td>46</td>
			<td>2018-2019-1</td>
			<td align="left">06220610</td>
			<td align="left">数据库原理</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000367&zcj=84',700,500)">84</a></td>


			<td>3</td>
			<td>56</td>
			<td>考试</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>47</td>
			<td>2018-2019-1</td>
			<td align="left">06230110</td>
			<td align="left">单片机及应用</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000368&zcj=82',700,500)">82</a></td>


			<td>3</td>
			<td>56</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>48</td>
			<td>2018-2019-1</td>
			<td align="left">06240210</td>
			<td align="left">面向对象的技术与UML</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000374&zcj=63',700,500)">63</a></td>


			<td>3.5</td>
			<td>64</td>
			<td>考试</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>49</td>
			<td>2018-2019-1</td>
			<td align="left">06240310</td>
			<td align="left">软件测试与项目管理</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000372&zcj=81',700,500)">81</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考试</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>50</td>
			<td>2018-2019-1</td>
			<td align="left">06262210</td>
			<td align="left">.Net平台与C#编程</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191000369&zcj=86',700,500)">86</a></td>


			<td>3</td>
			<td>56</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>51</td>
			<td>2018-2019-1</td>
			<td align="left">06291410</td>
			<td align="left">数据库应用课程设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820191001398&zcj=92',700,500)">92</a></td>


			<td>1</td>
			<td>1</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>52</td>
			<td>2018-2019-2</td>
			<td align="left">06215310</td>
			<td align="left">科技英语</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192001074&zcj=83',700,500)">83</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>53</td>
			<td>2018-2019-2</td>
			<td align="left">06240410</td>
			<td align="left">软件系统设计与体系结构</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192001075&zcj=80',700,500)">80</a></td>


			<td>3</td>
			<td>64</td>
			<td>考查</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>54</td>
			<td>2018-2019-2</td>
			<td align="left">06240510</td>
			<td align="left">XML技术</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192002087&zcj=87',700,500)">87</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考查</td>
			<td>必修</td>
			<td>专业课程</td>
		</tr>

		<tr>
			<td>55</td>
			<td>2018-2019-2</td>
			<td align="left">06266010</td>
			<td align="left">嵌入式系统开发</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192001078&zcj=89',700,500)">89</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>56</td>
			<td>2018-2019-2</td>
			<td align="left">06268310</td>
			<td align="left">Oracle网络数据库编程</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192001083&zcj=80',700,500)">80</a></td>


			<td>2.5</td>
			<td>48</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>57</td>
			<td>2018-2019-2</td>
			<td align="left">06291910</td>
			<td align="left">新技术与学科创新</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201820192001073&zcj=80',700,500)">80</a></td>


			<td>2</td>
			<td>2</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>58</td>
			<td>2019-2020-1</td>
			<td align="left">06216410</td>
			<td align="left">物联网技术</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201000921&zcj=86',700,500)">86</a></td>


			<td>2</td>
			<td>32</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>59</td>
			<td>2019-2020-1</td>
			<td align="left">06263310</td>
			<td align="left">Android应用开发</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201000927&zcj=88',700,500)">88</a></td>


			<td>3</td>
			<td>56</td>
			<td>考查</td>
			<td>限选</td>
			<td>专业选修课程</td>
		</tr>

		<tr>
			<td>60</td>
			<td>2019-2020-1</td>
			<td align="left">06290110</td>
			<td align="left">毕业实习</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201000067&zcj=86',700,500)">86</a></td>


			<td>4</td>
			<td>4</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>61</td>
			<td>2019-2020-1</td>
			<td align="left">06290210</td>
			<td align="left">毕业见习</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201000920&zcj=83',700,500)">83</a></td>


			<td>1</td>
			<td>1</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>62</td>
			<td>2019-2020-1</td>
			<td align="left">06292010</td>
			<td align="left">专业综合设计</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201000922&zcj=82',700,500)">82</a></td>


			<td>3</td>
			<td>3</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>63</td>
			<td>2019-2020-1</td>
			<td align="left">20200110</td>
			<td align="left">就业指导</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920201003966&zcj=通过',700,500)">通过</a></td>


			<td>1</td>
			<td>1</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

		<tr>
			<td>64</td>
			<td>2019-2020-2</td>
			<td align="left">06290010</td>
			<td align="left">毕业设计（论文）</td>

			<td style=" "><a href="javascript:JsMod('/jsxsd/kscj/pscj_list.do?xs0101id=201601013109&jx0404id=201920202001860&zcj=86',700,500)">86</a></td>


			<td>12</td>
			<td>12</td>
			<td>考查</td>
			<td>必修</td>
			<td>实践课</td>
		</tr>

	</table>
</div>
<br />

<div id="Footer1_divCopyright" class="Nsb_pw" style="display:;">
  <div class="Nsb_rights">Copyright (C) 湖南强智科技发展有限公司 2003-2013 All Rights Reserved 湘ICP 备12010071号</div>
</div>
<script language="javascript">

//合并学年学期列
//jQuery("#dataList").rowspan(1);
</script>
</body>
</html>`

func parseHtml(htmlText string) ([]string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlText))
	if err != nil {
		return nil, err
	}

	infoList := make([]string, 0)
	dom.Find("tbody>tr>td").Each(func(i int, selection *goquery.Selection) {
		str := strings.TrimSpace(selection.Text())
		if str != "" && str != " " {
			infoList = append(infoList, str)
		}
	})
	return infoList, nil
}

//  1 2016-2017-1 01100110 地球科学概论 72 4 64 考试 必修 学科基础课程
type itemScore struct {
	ID               string
	OpenDate         string
	CourseID         string
	CourseName       string
	Score            string
	ScoreCredit      string
	ScoreTime        string
	AssessmentMethod string
	AssessmentAttr   string
	CourseNature     string
}

func buildComplete(source []string) []*itemScore {
	length := len(source)
	if length%10 != 0 {
		fmt.Println("end", length, source[0])
		return nil
	}

	buildResult := make([]*itemScore, 0)
	for i := 0; i < length; i += 10 {
		item := &itemScore{
			ID:               source[i],
			OpenDate:         source[i+1],
			CourseID:         source[i+2],
			CourseName:       source[i+3],
			Score:            source[i+4],
			ScoreCredit:      source[i+5],
			ScoreTime:        source[i+6],
			AssessmentMethod: source[i+7],
			AssessmentAttr:   source[i+8],
			CourseNature:     source[i+9],
		}
		buildResult = append(buildResult, item)
	}
	return buildResult
}

func main() {
	res, _ := parseHtml(html_text)
	buildComplete(res)
}
