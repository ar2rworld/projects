import type { AxiosError } from 'axios';
import { LoginTokens } from '../stores/loginTokens';
import type { ILoginTokens } from '../types/login';
import type { IMe } from '../types/me';
import { api } from './axios';

export const getMe = () => {
	let lt: ILoginTokens = { accessToken: '', refreshToken: '', expiresIn: 0 };
	LoginTokens.update((v) => {
		lt = v;
		return v;
	});

	const headers: HeadersInit = {
		Authorization: `Bearer ${lt.accessToken}`
	};

	let out: IMe = {
		Ok: false,
		Error: '',
		Message: '',
		Username: '',
		FirstName: '',
		LastName: '',
		FullName: '',
		EmailVerified: false
	};

	api
		.get('/me', { headers: headers })
		.then((d) => {
			const me = d.data as IMe;

			out = {
				Ok: true,
				Error: '',
				Message: '',
				Username: me.Username,
				FirstName: me.FirstName,
				LastName: me.LastName,
				FullName: me.FullName,
				EmailVerified: me.EmailVerified
			};
		})
		.catch((e: AxiosError) => {
			out = e.response?.data as IMe;
		});

	return out;
};
