export interface PopupButton {
    id?: string;
    type?: 'default' | 'ok' | 'close' | 'cancel' | 'destructive';
    text?: string;
}

export interface PopupParams {
    title?: string;
    message: string;
    buttons?: PopupButton[];
}

export interface WebAppUser {
    id: number;
    first_name: string;
    last_name?: string;
    photo_url?: string;
}

interface WebAppInitData {
    user: WebAppUser
}

interface TelegramWebApp {
    initData: string;
    initDataUnsafe: WebAppInitData;
    ready(): void;
    showAlert(message: string, callback?: () => void): void;
    showConfirm(message: string, callback?: (result: boolean) => void): void;
    showPopup(params: PopupParams, callback?: (result: string) => void): void;
    setHeaderColor(color: string): void;
    setBackgroundColor(color: string): void;
}

interface TelegramWindow extends Window {
    Telegram?: {
        WebApp: TelegramWebApp;
    };
}

export const telegram = (window as unknown as TelegramWindow).Telegram!.WebApp
export const currentUser = telegram!.initDataUnsafe.user;