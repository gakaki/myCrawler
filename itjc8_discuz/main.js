import { chromium } from "playwright-core";
import child_process from "child_process";

async function launch(executablePath: string) {
  let browser;
  var connect = () => chromium.connectOverCDP("http://localhost:9222");
  try {
    browser = await connect();
  } catch (e) {
    child_process.exec(`"${executablePath}" --remote-debugging-port=9222`);
    await new Promise((r) => setTimeout(r, 5000));
    browser = await connect();
  }
  return browser;
}

launch(`C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe`)
  .then(async (browser) => {
    const [ctx] = browser.contexts();
    const [page] = ctx.pages();
    await page.goto("https://www.baidu.com/");
    await page.pause();
  })
  .catch(console.error);


