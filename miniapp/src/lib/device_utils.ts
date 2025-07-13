import {UAParser} from "ua-parser-js";

export type Device = 'mobile' | 'tablet' | 'desktop';

export const maxWidth = {
    mobile: 440,
    tablet: 1020
};

export interface UserAgent {
    os: 'ios' | 'ipad' | 'android' | 'macos' | 'windows' | 'linux' | undefined;
    browser: 'chrome' | 'firefox' | 'safari' | 'opera' | undefined;
}

export function isDevice(device: keyof typeof maxWidth | 'desktop'): boolean {
    if (!window) {
        return device === 'desktop';
    }

    const width = window.innerWidth;
    switch (device) {
        case 'desktop':
            return width > maxWidth.tablet;
        case 'tablet':
            return width > maxWidth.mobile;
        default:
        case 'mobile':
            return width <= maxWidth.mobile || isOS('ios', 'android', 'ipad');
    }
}

export function getUserAgent(): UserAgent {
    const results = new UAParser().getResult();
    const osName = results.os.name?.toLowerCase();
    const deviceModel = results.device.model?.toLowerCase();
    let os: UserAgent['os'];
    switch (true) {
        case deviceModel === 'ipad':
            os = 'ipad';
            break;
        case osName === 'ios':
            os = 'ios';
            break;
        case osName === 'android':
            os = 'android';
            break;
        case osName === 'mac os':
            os = 'macos';
            break;
        case osName === 'linux':
            os = 'linux';
            break;
        case osName?.includes('windows'):
            os = 'windows';
            break;
    }

    const browserName = results.browser.name?.toLowerCase();
    let browser: UserAgent['browser'] | undefined;
    switch (true) {
        case browserName === 'chrome':
            browser = 'chrome';
            break;
        case browserName === 'firefox':
            browser = 'firefox';
            break;
        case browserName?.includes('safari'):
            browser = 'safari';
            break;
        case browserName?.includes('opera'):
            browser = 'opera';
            break;
    }

    return {
        os,
        browser
    };
}

export function isOS(...os: UserAgent['os'][]): boolean {
    return os.includes(getUserAgent().os);
}

export const isDesktop = isOS('macos', 'windows', 'linux');
export const isTgUI = isOS('android') || isDesktop;