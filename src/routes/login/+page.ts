export function load(data) {
	return {
		new: data.url.searchParams.has("new")
	};
}
