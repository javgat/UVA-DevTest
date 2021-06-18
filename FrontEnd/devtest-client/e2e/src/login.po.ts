import { browser, by, element } from 'protractor';

export class LoginPage {

  getUrl(): string{
    return browser.baseUrl+"login"
  }

  async navigateTo(): Promise<unknown> {
    return browser.get(this.getUrl());
  }

  async setLoginid(loginid: string): Promise<void> {
    return element(by.name("inputUsername")).sendKeys(loginid)
  }

  async setPassword(password: string): Promise<void> {
    return element(by.name("inputPassword")).sendKeys(password)
  }

  async pressSubmit(): Promise<void> {
    return element(by.id("buttonLoginSignin")).click()
  }

  async loginUser(loginid: string, password: string){
    await this.setLoginid(loginid)
    await this.setPassword(password)
    return this.pressSubmit()
  }
}