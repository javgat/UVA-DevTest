import { browser, by, element } from 'protractor';

export class SigninPage {

  getUrl(): string{
    return browser.baseUrl+"signin"
  }

  async navigateTo(): Promise<unknown> {
    return browser.get(this.getUrl());
  }

  async setUsername(username: string): Promise<void> {
    return element(by.name("inputUsername")).sendKeys(username)
  }

  async setEmail(email: string): Promise<void> {
    return element(by.name("inputEmail")).sendKeys(email)
  }

  async setPassword(password: string): Promise<void> {
    return element(by.name("inputPassword")).sendKeys(password)
  }

  async registerUser(username: string, email: string, password: string): Promise<void>{
    await this.setUsername(username)
    await this.setEmail(email)
    await this.setPassword(password)
    return this.pressSubmit()
  }

  async pressSubmit(): Promise<void> {
    return element(by.id("buttonSubmitSignin")).click()
  }

  async getTitleText(): Promise<string> {
    return element(by.css('app-root .content span')).getText();
  }

}
