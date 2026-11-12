# Google Maps API用法教程

在过去的一年中，在线地图的发展是相当巨大，我们可以看到在线地图的极有价值的信息和其能力。这其中，最有名气的自然是Google Maps。. Google Maps由一个相当强大的开发引擎并也有一个很大的社区提示支持。

Google 允许各种web masters 通过Google Maps API来增加或自定义他们站点特定的地图，你可能从这里取得[Google API key](https://code.google.com/intl/en/apis/maps/signup.html "Get a Google Maps API Key") 。一个地图 API key只对一个“目录”或域有效。key绑定了你的域名，你要在网站上放地图，需要有对应的key，否则拒绝读取地图数据，在本地测试可以不用key。当然，你可以申请多个API key。

目录


- [创建一个简单的地图](#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E7%9A%84%E5%9C%B0%E5%9B%BE "创建一个简单的地图")
- [地图标记](#%E5%9C%B0%E5%9B%BE%E6%A0%87%E8%AE%B0 "地图标记")
- [活动标记](#%E6%B4%BB%E5%8A%A8%E6%A0%87%E8%AE%B0 "活动标记")
- [加入多个标记并分组](#%E5%8A%A0%E5%85%A5%E5%A4%9A%E4%B8%AA%E6%A0%87%E8%AE%B0%E5%B9%B6%E5%88%86%E7%BB%84 "加入多个标记并分组")
- [过滤显示标记](#%E8%BF%87%E6%BB%A4%E6%98%BE%E7%A4%BA%E6%A0%87%E8%AE%B0 "过滤显示标记")
#### 创建一个简单的地图

在你的站点上引入Google Maps 是一件很简单的事情，你只需要加入：

- 引入Google的JavaScript 文件
- 设置JavaScript 一些参数
- 一个你需要显示地图的HTML layer

**Google的Javascript文件引入**:

```js
<script
    charset="UTF-8"
    src=http://maps.google.com/maps?file=api&v=2&hl=en&oe=utf-8&key=API_KEY
    type="text/javascript">
</script>

```

***注意**：* 我们可以改变语言，比如说，把“**hl=en**” 改成中文“**hl=zh-CN**” 。我们还得要把“**API\_KEY**” 改成我们向Google申请来的那个。

**说明**: 使用 UTF-8 编码会更好些。

**设置地图参数**:

这是你可定制有个性化的Google Maps的地方。我们可以增加一些参数来改变地图的样式。例如，我们可以设置地图的载入和显示的坐标。下面是相关的代码：

```js
]function initialize() {
    if (GBrowserIsCompatible()) {
        var map = new GMap2(document.getElementById("map_canvas"));
        map.setCenter(new GLatLng(37.97918, 23.71665), 13);
        map.setUIToDefault();
    }
}
```

请注意上面高亮的那一条语句，第一个是纬度坐标和第二个是经度坐标，“13” 表示地图缩放的程度，这个值可以取1 到17。

要知道所在地点的纬度和经度，你可以使用[这个工具](http://www.satsig.net/maps/lat-long-finder.htm "Tool to get the coordinates of a location")，这个工具很容易使用，只需要把地图移到你想要的区域，然后，把鼠标放在中心就可以了。

#### 地图标记

你可以在地图上放上一个标记来标出一个特定的位置，下面是一个示例代码。

```js
var point = new GLatLng(37.97110, 23.72601);
map.addOverlay(new GMarker(point));
```

于是，我们整个代码看起来是下面这个样子：

```js
function initialize() {
    if (GBrowserIsCompatible()) {
        var map = new GMap2(document.getElementById("map_canvas"));
        map.setCenter(new GLatLng(37.97918, 23.71665), 13);
        var point = new GLatLng(37.97110, 23.72601);
        map.addOverlay(new GMarker(point));
        map.setUIToDefault();
    }
}
```

上面的示例把我们的地图的中心放在了希腊雅典，标记了雅典卫城。

**气球提示**

气球提示一个很不错的界面，他可以用于放置一些小提示或或是一些信息。例如，下面的代码将放置一个在雅典卫城山上放一个气球提示来显示一些信息：

```js
function initialize() {
    if (GBrowserIsCompatible()) {
        var map = new GMap2(document.getElementById("map_canvas"));
        map.setCenter(new GLatLng(37.97110, 23.72601), 13);
        var html = "Parthenon 帕台农神庙, 地址: Acropolis Hill";
        map.openInfoWindow(map.getCenter(),
        document.createTextNode(html));
        map.setUIToDefault();
    }
}
```
#### 活动标记

我们可以合并一些标记和气球提示来创建一个活动标记，这样一来，我们可以达到这样的效果——当用户点击某个标记的时候才会显示提示。代码如下所示：

```js
function initialize() {
    if (GBrowserIsCompatible()) {
        var map = new GMap2(document.getElementById("map_canvas"));
        map.setCenter(new GLatLng(37.97918, 23.71665), 13);
        var baseIcon = new GIcon(G_DEFAULT_ICON);
        baseIcon.shadow = "http://www.google.com/mapfiles/shadow50.png";
        baseIcon.iconSize = new GSize(20, 34);
        baseIcon.shadowSize = new GSize(37, 34);
        baseIcon.iconAnchor = new GPoint(9, 34);
        baseIcon.infoWindowAnchor = new GPoint(9, 2);

        function createMarker(point, index) {
            var redIcon = new GIcon(baseIcon);
            redIcon.image = "http://www.google.com/mapfiles/marker.png";
            markerOptions = { icon:redIcon };
            var marker = new GMarker(point, markerOptions);
            GEvent.addListener(marker, "click", function() {
                marker.openInfoWindowHtml("Parthenon, Address: Acropolis Hill");
            });
            return marker;
        }

        var latlng = new GLatLng(37.97110, 23.72601);
        map.addOverlay(createMarker(latlng));
    }
}
```

让我来梳理一下上面的代码。下面的部分是在标记下增加一个阴影：

```js
var baseIcon = new GIcon(G_DEFAULT_ICON);
baseIcon.shadow = "http://www.google.com/mapfiles/shadow50.png";
baseIcon.iconSize = new GSize(20, 34);
baseIcon.shadowSize = new GSize(37, 34);
baseIcon.iconAnchor = new GPoint(9, 34);
baseIcon.infoWindowAnchor = new GPoint(9, 2);
```

标记的Action是在这里设置的：

```js
function createMarker(point, index) {
    var redIcon = new GIcon(baseIcon);
    redIcon.image = "http://www.google.com/mapfiles/marker.png";
    markerOptions = { icon:redIcon };
    var marker = new GMarker(point, markerOptions);
    GEvent.addListener(marker, "click", function() {
        marker.openInfoWindowHtml("Parthenon, Address: Acropolis Hill");
    });
    return marker;
}
```

这里是我们的标记的坐标：

```js
var latlng = new GLatLng(37.97110, 23.72601);
map.addOverlay(createMarker(latlng));
```
### 载入地图

我们可以通过两种方法载入地图。我们可以让地图在整网页载入时载入（通过使用body的load事件），也可以把载入过程赋给其它事件。下面的两个方法是我们需要注意的：

- **initialize()** 载入地图
- **GUnload()** 卸载地图以释放内存

我们当然还需要使用HTML的DIV标签来指定一个ID，这样才能被JavaScript使用，在我们的示例中，我们使用“map\_canvas”。我们也能使用CSS来渲染这个DIV层。

### 定制地图

你可以使用自定义的标记和阴影。有两个工具可以帮助你创建这些图标—— [地图标记](https://gmaps-utility-library.googlecode.com/svn/trunk/mapiconmaker/1.1/examples/markericonoptions-wizard.html "Custom Marker Icons") 和 [阴影](http://www.cycloloco.com/shadowmaker/ "Custom Shadows for Maps")。你也可以使用HTML和CSS来定义气球提示。

#### 加入多个标记并分组

我们可以标记多个地点，并可以把它们根据我们的需要分组。这样一来，我们可以通过不同的标记图标来显示地点的不同属性，比如：酒店，车站等。要做到这样，我们只需要使用一个XML文件，一个简单的XML文件如下所示：

```xml
<markers>
<marker
    name="标题"
    label="这是一个标签"
    desc="气球提示的描述"
    lat="37.97167" lng="23.72581"
    type="标签的种类，如 Bridge"
    address="地址信息"
    icona="图标"
/>
</markers>

```

你可以在这个XML文件中加入多个地点信息。有一件事你需要小心的是，当出现一一些特定字符时（比如单引号，双引号，[“#@$](https://coolshell.cn/cdn-cgi/l/email-protection#17f5978b345733)<>”等），你需要使用HTML的转义。

使用这XML的脚本如下：

`<script src="http://gmaps-utility-library.googlecode.com/svn/trunk/labeledmarker/release/src/labeledmarker.js" type="text/javascript"></script>`

当然，你需要设置一些参数：

```js
var iconRed = new GIcon();
iconRed.image = '../img/marker-red.png';
iconRed.shadow = '';
iconRed.iconSize = new GSize(32, 32);
iconRed.shadowSize = new GSize(22, 20);
iconRed.iconAnchor = new GPoint(16, 16);
iconRed.infoWindowAnchor = new GPoint(5, 1);
var customIcons = [];

customIcons["ancient"] = iconRed;
var markerGroups = { "ancient": []};
```

上面，我们给customIcons 的“ancient”属性设置成了iconRed ，于是我们应该在我们的XML文件中使用类型(ancient) ，如果我们把这个XML文件命名为： markers.xml，那么，我们的代码如下：

```js
GDownloadUrl("markers.xml", function(data) { //We tell Google Maps to load our file
        var xml = GXml.parse(data);
        var markers = xml.documentElement.getElementsByTagName("marker"); //and read markers
        for (var i = 0; i < markers.length; i++) {
            var name = markers[i].getAttribute("name"); //From here down we assign variables.
            var label = markers[i].getAttribute("label");
            var desc = markers[i].getAttribute("desc");
            var address = markers[i].getAttribute("address");
            var type = markers[i].getAttribute("type");
            var icona = markers[i].getAttribute("icona");
            var point = new GlatLng(parseFloat(markers[i].getAttribute("lat")), //and we set the lat-long
            parseFloat(markers[i].getAttribute("lng")));
            var marker = createMarker(point, name, label, desc, address, type, icona);
            map.addOverlay(marker);
        }
    });
}
}

function createMarker(point, name, label, desc, address, type, icona) {
    var marker = new LabeledMarker(point, {icon: customIcons[type], labelText: label, labelOffset: new GSize(-6, -8)})
};

```

要分组标记，你需要下面的代码：

```js
    markerGroups[type].push(marker);
    var html = "<img src=" + icona + " height=150 border=0 /><br><br><span><b>"+ name + "</b><br/>" +
            desc + "<br/><b>Address:</b> " + address + "<br/><br/><span>";
    GEvent.addListener(marker, 'click', function() {
        marker.openInfoWindowHtml(html);
    });
    return marker;
}

```

要使用不同的图标，你可以使用相同的方法。

```js
var iconRed = new GIcon();
iconRed.image = '../img/marker-red.png';
iconRed.shadow = '';
iconRed.iconSize = new GSize(32, 32);
iconRed.shadowSize = new GSize(22, 20);
iconRed.iconAnchor = new GPoint(16, 16);
iconRed.infoWindowAnchor = new GPoint(5, 1);

var iconGreen = new GIcon();
iconGreen.image = '../img/marker-green.png';
iconGreen.shadow = '';
iconGreen.iconSize = new GSize(32, 32);
iconGreen.shadowSize = new GSize(22, 20);
iconGreen.iconAnchor = new GPoint(16, 16);
iconGreen.infoWindowAnchor = new GPoint(5, 1);

var iconBrown = new GIcon();
iconBrown.image = '../img/marker-brown.png';
iconBrown.shadow = '';
iconBrown.iconSize = new GSize(32, 32);
iconBrown.shadowSize = new GSize(22, 20);
iconBrown.iconAnchor = new GPoint(16, 16);
iconBrown.infoWindowAnchor = new GPoint(5, 1);

var customIcons = [];

customIcons["hotel"] = iconRed;
customIcons["bridge"] = iconGreen;
customIcons["hill"] = iconBrown;
var markerGroups = { "hotel": [], "bridge": [], "hill": []};
```

所以，如果我们在XML 文件中设置了不同的种类，如：hotel，bridge 和 hill，我们也应该需要不同的颜色和图标。

#### 过滤显示标记

我们还可以让我们的标记更友好一些。我们可以让用户决定是否显示标记，这样的话，地图上的标记就不会太多，也会根据用户的需求来显示相当的标记。我们可以使用几个按钮，复选框，或是链接来完成这个事情。要做到这个事，你需要在“*map.addMapType(G\_SATELLITE\_3D\_MAP);* ”后面加入下面的代码：

```js
document.getElementById("hotelCheckbox").checked = true;
document.getElementById("bridgeCheckbox").checked = true;
document.getElementById("hillCheckbox").checked = true;
document.getElementById("labelsCheckbox").checked = true;

```

然后再加入下面的这些 JavaScript 的代码：

```js
function toggleGroup(type) {
    for (var i = 0; i < markerGroups[type].length; i++) {
        var marker = markerGroups[type][i];
        if (marker.isHidden()) {
            marker.show();
        } else {
            marker.hide();
        }
    }
}

function toggleLabels() {
    var showLabels = document.getElementById("labelsCheckbox").checked;
    for (groupName in markerGroups) {
        for (var i = 0; i < markerGroups[groupName].length; i++) {
            var marker = markerGroups[groupName][i];
            marker.setLabelVisibility(showLabels);
        }
    }
}

function hideAll() {
    var boxes = document.getElementsByName("mark");
    for(var i = 0; i < boxes.length; i++) {
        if(boxes[i].checked) {
            boxes[i].checked = false;
            switchLayer(false, layers[i].obj);
            chosen.push(i);
        }
    }
}

function checkChecked() {
    var boxes = document.getElementsByName("mark");
    for(var i = 0; i < boxes.length; i++) {
        if(boxes[i].checked) return true;
    }
    return false;
}
```

最后一件事是加如几个checkbox ：

```xml
<input type="hidden" id="labelsCheckbox" onclick="toggleLabels()" checked=”checked” />
<label for=”hotelCheckbox”>Hotels</label><input type="checkbox" id="hotelCheckbox" onclick="toggleGroup('hotel')" checked=”checked” />
<label for=”bridgeCheckbox”>Bridges</label><input type="checkbox" id="bridgeCheckbox" onclick="toggleGroup('bridge')" checked=”checked” />

```

我们 Google Maps 就绪了，这篇文章讲述了Google Map API中你应该知道的。希望这篇文章对你有帮助。

文章：[来源](http://jeez.eu/2009/10/09/google-maps-from-a-to-z/)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
