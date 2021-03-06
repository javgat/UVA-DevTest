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

import { CustomizedView } from '../model/customizedView';
import { EmailConfiguration } from '../model/emailConfiguration';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class ConfigurationService {

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
     * Returns a CustomizedView
     * Returns a CustomizedView
     * @param rolBase Name of the rolBase to find its customized view
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', observe?: 'body', reportProgress?: boolean): Observable<CustomizedView>;
    public getCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<CustomizedView>>;
    public getCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<CustomizedView>>;
    public getCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (rolBase === null || rolBase === undefined) {
            throw new Error('Required parameter rolBase was null or undefined when calling getCView.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // authentication (NoRegistered) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["NotLoggedIn"]) {
            headers = headers.set('NotLoggedIn', this.configuration.apiKeys["NotLoggedIn"]);
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
        ];

        return this.httpClient.get<CustomizedView>(`${this.basePath}/customizedViews/${encodeURIComponent(String(rolBase))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns all CustomizedViews
     * Returns all CustomizedViews
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getCViews(observe?: 'body', reportProgress?: boolean): Observable<Array<CustomizedView>>;
    public getCViews(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<CustomizedView>>>;
    public getCViews(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<CustomizedView>>>;
    public getCViews(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // authentication (NoRegistered) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["NotLoggedIn"]) {
            headers = headers.set('NotLoggedIn', this.configuration.apiKeys["NotLoggedIn"]);
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
        ];

        return this.httpClient.get<Array<CustomizedView>>(`${this.basePath}/customizedViews`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns the information about the current configuration related to email service
     * Returns the information about the current configuration related to email service
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getEmailConfiguration(observe?: 'body', reportProgress?: boolean): Observable<EmailConfiguration>;
    public getEmailConfiguration(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<EmailConfiguration>>;
    public getEmailConfiguration(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<EmailConfiguration>>;
    public getEmailConfiguration(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
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
        ];

        return this.httpClient.get<EmailConfiguration>(`${this.basePath}/emailConfiguration`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Modifies a CustomizedView
     * Modifies a CustomizedView
     * @param rolBase Name of the rolBase to modify its customized view
     * @param newView Customized View information updated
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', newView: CustomizedView, observe?: 'body', reportProgress?: boolean): Observable<CustomizedView>;
    public putCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', newView: CustomizedView, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<CustomizedView>>;
    public putCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', newView: CustomizedView, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<CustomizedView>>;
    public putCView(rolBase: 'administrador' | 'profesor' | 'estudiante' | 'noRegistrado', newView: CustomizedView, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (rolBase === null || rolBase === undefined) {
            throw new Error('Required parameter rolBase was null or undefined when calling putCView.');
        }

        if (newView === null || newView === undefined) {
            throw new Error('Required parameter newView was null or undefined when calling putCView.');
        }

        let headers = this.defaultHeaders;

        // authentication (BearerCookie) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["Cookie"]) {
            headers = headers.set('Cookie', this.configuration.apiKeys["Cookie"]);
        }

        // authentication (NoRegistered) required
        if (this.configuration.apiKeys && this.configuration.apiKeys["NotLoggedIn"]) {
            headers = headers.set('NotLoggedIn', this.configuration.apiKeys["NotLoggedIn"]);
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

        return this.httpClient.put<CustomizedView>(`${this.basePath}/customizedViews/${encodeURIComponent(String(rolBase))}`,
            newView,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Modifies the information related to the configuration of the email service
     * Modifies the information related to the configuration of the email service
     * @param emailConfiguration New email configuration
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putEmailConfiguration(emailConfiguration: EmailConfiguration, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public putEmailConfiguration(emailConfiguration: EmailConfiguration, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public putEmailConfiguration(emailConfiguration: EmailConfiguration, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public putEmailConfiguration(emailConfiguration: EmailConfiguration, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (emailConfiguration === null || emailConfiguration === undefined) {
            throw new Error('Required parameter emailConfiguration was null or undefined when calling putEmailConfiguration.');
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

        return this.httpClient.put<any>(`${this.basePath}/emailConfiguration`,
            emailConfiguration,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}
