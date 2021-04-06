/**
 * DevTest
 * DevTest
 *
 * OpenAPI spec version: 1.0.0
 * Contact: javigaton@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */
/* tslint:disable:no-unused-variable member-ordering */

import { Inject, Injectable, Optional }                      from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams,
         HttpResponse, HttpEvent }                           from '@angular/common/http';
import { CustomHttpUrlEncodingCodec }                        from '../encoder';

import { Observable }                                        from 'rxjs';

import { LoginUser } from '../model/loginUser';
import { Password } from '../model/password';
import { PasswordUpdate } from '../model/passwordUpdate';
import { SigninUser } from '../model/signinUser';
import { User } from '../model/user';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class AuthService {

    protected basePath = 'https://localhost/DevTest';
    public defaultHeaders = new HttpHeaders();
    public configuration = new Configuration();

    constructor(protected httpClient: HttpClient, @Optional()@Inject(BASE_PATH) basePath: string, @Optional() configuration: Configuration) {
        if (basePath) {
            this.basePath = basePath;
        }
        if (configuration) {
            this.configuration = configuration;
            this.basePath = basePath || configuration.basePath || this.basePath;
        }
    }

    /**
     * @param consumes string[] mime-types
     * @return true: consumes contains 'multipart/form-data', false: otherwise
     */
    private canConsumeForm(consumes: string[]): boolean {
        const form = 'multipart/form-data';
        for (const consume of consumes) {
            if (form === consume) {
                return true;
            }
        }
        return false;
    }


    /**
     * Deletes all sessions of the user. Makes every current JWT related to him useless.
     * Deletes all sessions of the user. Makes every current JWT related to him useless.
     * @param username Username of the user with the token
     * @param password Current password of the user
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public closeSessions(username: string, password: Password, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public closeSessions(username: string, password: Password, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public closeSessions(username: string, password: Password, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public closeSessions(username: string, password: Password, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (username === null || username === undefined) {
            throw new Error('Required parameter username was null or undefined when calling closeSessions.');
        }

        if (password === null || password === undefined) {
            throw new Error('Required parameter password was null or undefined when calling closeSessions.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.delete<any>(`${this.basePath}/accesstokens/${encodeURIComponent(String(username))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Registers a new authorized connection token
     * Tries to login, and gets a JWT auth token if successful
     * @param loginUser User who is trying to generate a token
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public login(loginUser: LoginUser, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public login(loginUser: LoginUser, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public login(loginUser: LoginUser, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public login(loginUser: LoginUser, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (loginUser === null || loginUser === undefined) {
            throw new Error('Required parameter loginUser was null or undefined when calling login.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.post<any>(`${this.basePath}/accesstokens`,
            loginUser,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns a useless cookie that will expire soon
     * Returns a useless cookie that will expire soon
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public logout(observe?: 'body', reportProgress?: boolean): Observable<any>;
    public logout(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public logout(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public logout(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.get<any>(`${this.basePath}/logout`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Modifies the password of the user &lt;username&gt;
     * Modifies the password of the user &lt;username&gt;
     * @param username Username of the user to modify its password
     * @param passwordUpdate Password update information
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putPassword(username: string, passwordUpdate: PasswordUpdate, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public putPassword(username: string, passwordUpdate: PasswordUpdate, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public putPassword(username: string, passwordUpdate: PasswordUpdate, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public putPassword(username: string, passwordUpdate: PasswordUpdate, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (username === null || username === undefined) {
            throw new Error('Required parameter username was null or undefined when calling putPassword.');
        }

        if (passwordUpdate === null || passwordUpdate === undefined) {
            throw new Error('Required parameter passwordUpdate was null or undefined when calling putPassword.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.put<any>(`${this.basePath}/users/${encodeURIComponent(String(username))}/password`,
            passwordUpdate,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * adds a user
     * Adds a user to the system
     * @param signinUser User item to add
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public registerUser(signinUser: SigninUser, observe?: 'body', reportProgress?: boolean): Observable<User>;
    public registerUser(signinUser: SigninUser, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<User>>;
    public registerUser(signinUser: SigninUser, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<User>>;
    public registerUser(signinUser: SigninUser, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (signinUser === null || signinUser === undefined) {
            throw new Error('Required parameter signinUser was null or undefined when calling registerUser.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.post<User>(`${this.basePath}/users`,
            signinUser,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Modifies the current JWT Cookie related to the current session, extending it.
     * Modifies the current JWT Cookie related to the current session, extending.
     * @param username Username of the user with the token
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public relogin(username: string, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public relogin(username: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public relogin(username: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public relogin(username: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (username === null || username === undefined) {
            throw new Error('Required parameter username was null or undefined when calling relogin.');
        }

        let headers = this.defaultHeaders;

        // authentication (ReAuthCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];

        return this.httpClient.put<any>(`${this.basePath}/accesstokens/${encodeURIComponent(String(username))}`,
            null,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}
