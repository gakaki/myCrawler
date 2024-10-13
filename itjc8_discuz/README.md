1pip install -u pytest-playwright
playwright install

# MacOS系统下：playwright如何连接已打开的谷歌浏览器，绕过反爬
    https://blog.csdn.net/weixin_56314697/article/details/137876565

export PATH="/Applications/Google Chrome.app/Contents/MacOS:$PATH"
Google\ Chrome --remote-debugging-port=9222


Google\ Chrome --remote-debugging-port=9222 --user-data-dir="~/ChromeProfile"


在终端中输入以下命令以使用文本编辑器打开.zshrc文件：
nano ~/.zshrc
 
在打开的.zshrc文件中找到一个空白行，或者在文件的末尾添加以下行：
export PATH="/Applications/Google Chrome.app/Contents/MacOS:$PATH"


保存对文件的更改。在nano编辑器中，按下Ctrl + O，然后按Enter键以保存文件。然后按下Ctrl + X以退出nano编辑器。

激活环境： 在终端中运行以下命令以使更改生效：
source ~/.zshrc
这将重新加载你的.zshrc文件，并使你添加的导出命令生效。


# exists playwright browser session
https://stackoverflow.com/questions/71362982/is-there-a-way-to-connect-to-my-existing-browser-session-using-playwright


#  playwright对已打开的浏览器进行连接、操作
~~
~~
## nodejs
https://github.com/microsoft/playwright/issues/11442








import time

from playwright.sync_api import Playwright,sync_playwright
# C:\Users\xiaozai\AppData\Local\ms-playwright
with sync_playwright() as playwright:
    browser = playwright.chromium.launch_persistent_context(
        # 指定本机用户缓存地址
        user_data_dir=r"C:\Users\xiaozai\AppData\Local\Google\Chrome\User Data",
        # 指定本机google客户端exe的路径
        executable_path=r"C:\Users\xiaozai\AppData\Local\Google\Chrome\Application\chrome.exe",
        # 要想通过这个下载文件这个必然要开  默认是False
        accept_downloads=True,
        # 设置不是无头模式
        headless=False,
        bypass_csp=True,
        slow_mo=10,
        # 跳过检测
        args = ['--disable-blink-features=AutomationControlled','--remote-debugging-port=9222']

    )
    page = browser.new_page()
    page.goto("https://www.baidu.com/")
    print(page.title())
    time.sleep(200)
    browser.close()
