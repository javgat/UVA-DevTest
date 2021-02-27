import { browser, logging } from 'protractor';
import { protractor } from 'protractor/built/ptor';
import { AppPage } from './app.po';
import { LoginPage } from './login.po';
import { SigninPage } from './signin.po';

class HelperFuncs{


  constructor(private page: AppPage, private signinPage: SigninPage, private loginPage: LoginPage){

  }

  async isErrorMessage(): Promise<boolean>{
    if(await this.page.messageVisible()){
      let type = await this.page.messageType()
      return type==="alert alert-danger"
    }
    return false
  }

  async isSuccessMessage(): Promise<boolean>{
    if(await this.page.messageVisible()){
      let type = await this.page.messageType()
      return type==="alert alert-success"
    }
    return false
  }

  async loginSuccess(): Promise<boolean>{
    return this.isSuccessMessage()
  }

  async loginWrong(): Promise<boolean>{
    return this.isErrorMessage()
  }
}

describe('workspace-project App', () => {
  let page: AppPage;
  let signinPage: SigninPage;
  let loginPage: LoginPage;
  let help: HelperFuncs
  let EC = protractor.ExpectedConditions

  beforeEach(() => {
    page = new AppPage();
    signinPage = new SigninPage()
    loginPage = new LoginPage()
    help = new HelperFuncs(page, signinPage, loginPage)
  });
  
  afterEach(async () => {
    //Esto no va en firefox he leido, por eso lo comento, pero no lo quito, investigare mas adelante
    /*
    // Assert that there are no errors emitted from the browser
    const logs = await browser.manage().logs().get(logging.Type.BROWSER);
    expect(logs).not.toContain(jasmine.objectContaining({
      level: logging.Level.SEVERE,
    } as logging.Entry));*/
  });

  it('should register user, log in as user, show errors when registering as existing user or logging in as unexisting user', async() => {
    let username = "username"
    let usernameWrong = "usernameWrong"
    let email = "username@email.com"
    let password = "password"
    // Register a new user
    await signinPage.navigateTo();
    await signinPage.registerUser(username, email, password)
    expect(await browser.wait(EC.urlContains(loginPage.getUrl()), 5000).catch(() => {return false}))
      .toBeTruthy("Redireccion de signin exitoso a login no se realizo")
    // Prompt an error when registering with an already registered user username/email data
    await signinPage.navigateTo()
    await signinPage.registerUser(username, email, password)
    expect(await help.isErrorMessage()).toBeTruthy("No hay mensaje de error al registrar usuario duplicado")
    await signinPage.registerUser(usernameWrong, email, password)
    expect(await help.isErrorMessage()).toBeTruthy("No hay mensaje de error al registrar email duplicado")
    // Log in as an existing user with username/email
    await loginPage.navigateTo()
    await loginPage.loginUser(username, password)
    expect(await help.loginSuccess()).toBeTruthy("No se detecta feedback de login funcional con username")
    await page.logout()
    await loginPage.navigateTo()
    await loginPage.loginUser(email, password)
    expect(await help.loginSuccess()).toBeTruthy("No se detecta feedback de login funcional con email")
    await page.logout()
    // Prompt error logging in as unexisting user 
    await loginPage.navigateTo()
    await loginPage.loginUser(usernameWrong, password)
    expect(await help.loginWrong()).toBeTruthy("No se detecta feedback de error de login con username inexistente")
  })
});
