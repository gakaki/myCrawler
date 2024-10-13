from playwright.sync_api import sync_playwright, Playwright

# # 使用 Playwright 连接到已经运行的 Chrome 浏览器
# with sync_playwright() as playwright:
#     # 连接已打开浏览器，找好端口
#     browser = playwright.chromium.connect_over_cdp('http://localhost:9222/')
# #     browser = playwright.chromium.connect_over_cdp("ws://127.0.0.1:9222/devtools/browser/f9e34e28-cf4c-4285-a21c-5f8f339dda1f")
#     default_context = browser.contexts[0]  # 注意这里不是browser.new_page()了
#     page = default_context.pages[0] #如果后续定位不到，说明page这里选择错了页面
#     print(page)
#
#     page.set_default_timeout(30000)


with sync_playwright() as p:
    browser = p.chromium.connect_over_cdp('http://localhost:9222/')
    # 获取page对象
    page = browser.contexts[0].pages[0]
    print(page.url)
    print(page.title())
    page.get_by_text('新随笔').click()