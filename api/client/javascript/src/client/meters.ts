import { transformResponse } from './utils.js'
import type { RequestOptions } from './common.js'
import type { MeterCreate, operations, paths } from './schemas.js'
import type { Client } from 'openapi-fetch'

/**
 * Meters
 * @description Meters are used to track and manage usage of your application.
 */
export class Meters {
  constructor(private client: Client<paths, `${string}/${string}`>) {}

  /**
   * Create a meter
   * @param meter - The meter to create
   * @param signal - An optional abort signal
   * @returns The created meter
   */
  public async create(meter: MeterCreate, options?: RequestOptions) {
    const resp = await this.client.POST('/api/v1/meters', {
      body: meter,
      ...options,
    })

    return transformResponse(resp)
  }

  /**
   * Get a meter by ID or slug
   * @param idOrSlug - The ID or slug of the meter
   * @param signal - An optional abort signal
   * @returns The meter
   */
  public async get(
    idOrSlug: operations['getMeter']['parameters']['path']['meterIdOrSlug'],
    options?: RequestOptions
  ) {
    const resp = await this.client.GET('/api/v1/meters/{meterIdOrSlug}', {
      params: {
        path: {
          meterIdOrSlug: idOrSlug,
        },
      },
      ...options,
    })

    return transformResponse(resp)
  }

  /**
   * List meters
   * @param signal - An optional abort signal
   * @returns The meters
   */
  public async list(options?: RequestOptions) {
    const resp = await this.client.GET('/api/v1/meters', {
      ...options,
    })

    return transformResponse(resp)
  }

  /**
   * Query usage data for a meter by ID or slug
   * @param idOrSlug - The ID or slug of the meter
   * @param query - The query parameters
   * @param signal - An optional abort signal
   * @returns The meter data
   */
  public async query(
    idOrSlug: operations['queryMeter']['parameters']['path']['meterIdOrSlug'],
    query?: operations['queryMeter']['parameters']['query'],
    options?: RequestOptions
  ) {
    const resp = await this.client.GET('/api/v1/meters/{meterIdOrSlug}/query', {
      params: {
        path: {
          meterIdOrSlug: idOrSlug,
        },
        query,
      },
      ...options,
    })

    return transformResponse(resp)
  }

  /**
   * Delete a meter by ID or slug
   * @param idOrSlug - The ID or slug of the meter
   * @param signal - An optional abort signal
   * @returns The deleted meter
   */
  public async delete(
    idOrSlug: operations['deleteMeter']['parameters']['path']['meterIdOrSlug'],
    options?: RequestOptions
  ) {
    const resp = await this.client.DELETE('/api/v1/meters/{meterIdOrSlug}', {
      params: {
        path: {
          meterIdOrSlug: idOrSlug,
        },
      },
      ...options,
    })

    return transformResponse(resp)
  }
}
