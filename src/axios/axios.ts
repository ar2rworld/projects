import axios from 'axios';

export const api = axios.create({
	baseURL: 'http://localhost:8050/projectsapi',
	timeout: 5000
});
