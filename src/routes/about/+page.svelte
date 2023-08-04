<script lang="ts">
	export let data;
	import { Table } from '@skeletonlabs/skeleton';
	import { tableMapperValues } from '@skeletonlabs/skeleton';
	$: elements = {
		// A list of heading labels.
		head: ['Name', 'Symbol', 'Weight'],
		// The data visibly shown in your table body UI.
		body: tableMapperValues(data.elements, ['name', 'symbol', 'weight']),
		// Optional: The data returned when interactive is enabled and a row is clicked.
		meta: tableMapperValues(data.elements, ['id', 'name', 'symbol', 'weight']),
		// Optional: A list of footer labels.
		foot: ['Total', '', `<code class="code">${data.elements.length}</code>`]
	};
	function rowClicked() {
		return;
	}
</script>

<Table source={elements} interactive={true} on:selected={rowClicked} />
