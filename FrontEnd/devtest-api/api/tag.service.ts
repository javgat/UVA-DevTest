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

import { Question } from '../model/question';
import { Tag } from '../model/tag';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class TagService {

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
     * Returns all questions from a tag.
     * Returns all questions from a tag.
     * @param tag Tag to find its questions
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getQuestionsFromTag(tag: string, observe?: 'body', reportProgress?: boolean): Observable<Array<Question>>;
    public getQuestionsFromTag(tag: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Question>>>;
    public getQuestionsFromTag(tag: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Question>>>;
    public getQuestionsFromTag(tag: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (tag === null || tag === undefined) {
            throw new Error('Required parameter tag was null or undefined when calling getQuestionsFromTag.');
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

        return this.httpClient.get<Array<Question>>(`${this.basePath}/tags/${encodeURIComponent(String(tag))}/questions`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns a tags.
     * Returns a tags.
     * @param tag Tag to find
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getTag(tag: string, observe?: 'body', reportProgress?: boolean): Observable<Tag>;
    public getTag(tag: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Tag>>;
    public getTag(tag: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Tag>>;
    public getTag(tag: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (tag === null || tag === undefined) {
            throw new Error('Required parameter tag was null or undefined when calling getTag.');
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

        return this.httpClient.get<Tag>(`${this.basePath}/tags/${encodeURIComponent(String(tag))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns all tags.
     * Returns all tags.
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getTags(observe?: 'body', reportProgress?: boolean): Observable<Array<Tag>>;
    public getTags(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Tag>>>;
    public getTags(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Tag>>>;
    public getTags(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

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

        return this.httpClient.get<Array<Tag>>(`${this.basePath}/tags`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}