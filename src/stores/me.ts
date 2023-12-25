import { writable } from 'svelte/store';
import type { IMe } from '../types/me';

const me: IMe = {
	Ok: false,
	Error: '',
	Message: '',
	Username: '',
	FirstName: '',
	LastName: '',
	FullName: '',
	EmailVerified: false
};

export const Me = writable(me);
