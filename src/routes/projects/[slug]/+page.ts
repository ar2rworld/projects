import { error } from '@sveltejs/kit';
import { projects } from '../data.js';

export function load({ params }) {
	const project = projects.find((project) => project.id === params.slug);

	if (!project) throw error(404);

	return {
		project
	};
}
