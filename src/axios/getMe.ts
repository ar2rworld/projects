import { AxiosError, HttpStatusCode } from 'axios';
import type { ILoginTokens } from '../types/login';
import type { IMe } from '../types/me';
import { api } from './axios';

export const getMe = async (lt: ILoginTokens): Promise<IMe> => {
	const headers: HeadersInit = {
		Authorization: `${lt.accessToken}`
	};

	const d = await api.get('/me', { headers: headers })

	const me = d.data as IMe;
	
	if (d.status != HttpStatusCode.Ok) {
		const err = d.data as AxiosError;
		me.Error = err.name
		me.Message = err.message
		return Promise.reject(me);
	}

	return Promise.resolve(me);
};
