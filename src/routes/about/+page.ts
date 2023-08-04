import { pb } from '$lib/pocketbase';
export const load = async ({ params }) => {
	const result = await pb.collection('elements').getFullList();
	console.log(result);
	return {
		elements: result
	};
};
