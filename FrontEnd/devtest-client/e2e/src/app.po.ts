import { browser, by, element } from 'protractor';

export class AppPage {

  getUrl(): string{
    return browser.baseUrl
  }

  async navigateTo(): Promise<unknown> {
    return browser.get(this.getUrl());
  }

  async messageVisible(): Promise<boolean>{
    return element(by.id("message")).isPresent()
  }

  async messageType(): Promise<string>{
    return element(by.id("message")).getAttribute("class")
  }

  async logout(){
    // Aqui hacer el logout, por ahora nada
  }
}
