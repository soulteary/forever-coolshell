# 如何使用Python操作摄像头

用过USB摄像头的都知道，你需要使用鼠标来操作它，比如截个图，录个像什么的，要点N次鼠标，对于我们那些不喜欢多次点击鼠标的人来说，这是一件很boring的事情，所以，本文将教你如何使用Python来操作摄像头。

这里，我们需要三个Python库： [VideoCapture](http://videocapture.sourceforge.net/)， [PIL](http://www.pythonware.com/products/pil/)  和 [pygame](http://www.pygame.org/)。使用这三个库你可以非常容易的编写一个摄像头程序。之所以使用pygame，其目的就是因为这个库可以处理视频帧（fps）。下面是代码：

```python
from VideoCapture import Device
import ImageDraw, sys, pygame, time
from pygame.locals import *
from PIL import ImageEnhance

res = (640,480)
pygame.init()
cam = Device()
cam.setResolution(res[0],res[1])
screen = pygame.display.set_mode((640,480))
pygame.display.set_caption('Webcam')
pygame.font.init()
font = pygame.font.SysFont("Courier",11)

def disp(phrase,loc):
    s = font.render(phrase, True, (200,200,200))
    sh = font.render(phrase, True, (50,50,50))
    screen.blit(sh, (loc[0]+1,loc[1]+1))
    screen.blit(s, loc)

brightness = 1.0
contrast = 1.0
shots = 0

while 1:
    camshot = ImageEnhance.Brightness(cam.getImage()).enhance(brightness)
    camshot = ImageEnhance.Contrast(camshot).enhance(contrast)
    for event in pygame.event.get():
        if event.type == pygame.QUIT: sys.exit()
    keyinput = pygame.key.get_pressed()
    if keyinput[K_1]: brightness -= .1
    if keyinput[K_2]: brightness += .1
    if keyinput[K_3]: contrast -= .1
    if keyinput[K_4]: contrast += .1
    if keyinput[K_q]: cam.displayCapturePinProperties()
    if keyinput[K_w]: cam.displayCaptureFilterProperties()
    if keyinput[K_s]:
        filename = str(time.time()) + ".jpg"
        cam.saveSnapshot(filename, quality=80, timestamp=0)
        shots += 1
    camshot = pygame.image.frombuffer(camshot.tostring(), res, "RGB")
    screen.blit(camshot, (0,0))
    disp("S:" + str(shots), (10,4))
    disp("B:" + str(brightness), (10,16))
    disp("C:" + str(contrast), (10,28))
    pygame.display.flip()
```

这段代码中的一些要点的解释如下：

- 第15行的那个函数是在视频上显示些信息。这个例子中，显示的是抓图的数量以及当前的亮度和对比度。这个函数先显示深灰色的文本，然后偏移几个像素，再显示浅灰色的，这样可以有阴影的效果。
- 第26行是在调整亮度和对比度。30-33行是在设置数字键1-4用于调整亮度和对比度。
- 34 和35行是在设置 ‘q’ 和 ‘w’ 来显示摄像头的对话框。在那里你可以调整分辨率和暴光度等等。
- 36行及以下的代码，是在存一个抓图文件。文件名中使用了当前时间。.

希望这个小程序能给你开启一个如何写摄像头的程序。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
