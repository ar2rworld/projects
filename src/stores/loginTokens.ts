import { writable } from 'svelte/store';
import type { ILoginTokens } from '../types/login';

const loginTokens: ILoginTokens = {
	accessToken: '',
	refreshToken: '',
	expiresIn: 0
};

export const LoginTokens = writable(loginTokens);
