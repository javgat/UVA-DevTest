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

import { TipoRol } from '../model/tipoRol';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class TiporolService {

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
     * Deletes a TipoRol.
     * Deletes a TipoRol. Every user with that TipoRol will now have the next less important TipoRol, or the next more important if there is none with less
     * @param rolNombre Nombre of the TipoRol to delete
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public deleteTipoRol(rolNombre: string, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public deleteTipoRol(rolNombre: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public deleteTipoRol(rolNombre: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public deleteTipoRol(rolNombre: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (rolNombre === null || rolNombre === undefined) {
            throw new Error('Required parameter rolNombre was null or undefined when calling deleteTipoRol.');
        }

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

        return this.httpClient.delete<any>(`${this.basePath}/tipoRoles/${encodeURIComponent(String(rolNombre))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns a TipoRol
     * Returns a TipoRol
     * @param rolNombre Nombre of the TipoRol to find
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getTipoRol(rolNombre: string, observe?: 'body', reportProgress?: boolean): Observable<TipoRol>;
    public getTipoRol(rolNombre: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<TipoRol>>;
    public getTipoRol(rolNombre: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<TipoRol>>;
    public getTipoRol(rolNombre: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (rolNombre === null || rolNombre === undefined) {
            throw new Error('Required parameter rolNombre was null or undefined when calling getTipoRol.');
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
        ];

        return this.httpClient.get<TipoRol>(`${this.basePath}/tipoRoles/${encodeURIComponent(String(rolNombre))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns all tipoRoles
     * Returns all tipoRoles
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getTipoRoles(observe?: 'body', reportProgress?: boolean): Observable<Array<TipoRol>>;
    public getTipoRoles(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<TipoRol>>>;
    public getTipoRoles(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<TipoRol>>>;
    public getTipoRoles(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

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

        return this.httpClient.get<Array<TipoRol>>(`${this.basePath}/tipoRoles`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Creates a TipoRol
     * Creates a TipoRol
     * @param newTipoRol New data for the TipoRol
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public postTipoRol(newTipoRol: TipoRol, observe?: 'body', reportProgress?: boolean): Observable<TipoRol>;
    public postTipoRol(newTipoRol: TipoRol, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<TipoRol>>;
    public postTipoRol(newTipoRol: TipoRol, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<TipoRol>>;
    public postTipoRol(newTipoRol: TipoRol, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (newTipoRol === null || newTipoRol === undefined) {
            throw new Error('Required parameter newTipoRol was null or undefined when calling postTipoRol.');
        }

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
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.post<TipoRol>(`${this.basePath}/tipoRoles`,
            newTipoRol,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Modifies a TipoRol
     * Modifies a TipoRol
     * @param rolNombre Nombre of the TipoRol to modify
     * @param newTipoRol New data for the TipoRol
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public putTipoRol(rolNombre: string, newTipoRol: TipoRol, observe?: 'body', reportProgress?: boolean): Observable<any>;
    public putTipoRol(rolNombre: string, newTipoRol: TipoRol, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<any>>;
    public putTipoRol(rolNombre: string, newTipoRol: TipoRol, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<any>>;
    public putTipoRol(rolNombre: string, newTipoRol: TipoRol, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (rolNombre === null || rolNombre === undefined) {
            throw new Error('Required parameter rolNombre was null or undefined when calling putTipoRol.');
        }

        if (newTipoRol === null || newTipoRol === undefined) {
            throw new Error('Required parameter newTipoRol was null or undefined when calling putTipoRol.');
        }

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
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.put<any>(`${this.basePath}/tipoRoles/${encodeURIComponent(String(rolNombre))}`,
            newTipoRol,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}