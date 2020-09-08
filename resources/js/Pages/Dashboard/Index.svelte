<script>
    import { InertiaLink, page } from '@inertiajs/inertia-svelte';
    import Layout from '@/Shared/Layout.svelte';
    import Helmet from '@/Shared/Helmet.svelte';
</script>

<Helmet title="Dashboard"/>

<Layout>
    <div>
        <h1 class="mb-8 font-bold text-3xl">Dashboard</h1>

        <p class="mb-6 leading-normal font-bold">
            Hey there!
        </p>

        <p class="mb-6 leading-normal">
            Welcome to Ping CRM, a demo app designed to help illustrate how
            <a class="text-indigo-600 underline hover:text-orange-500" href="https://inertiajs.com">Inertia.js</a>
            works with <a class="text-indigo-600 underline hover:text-orange-500" href="https://svelte.dev/">Svelte (client-side)</a>
            and <a class="text-indigo-600 underline hover:text-orange-500" href="https://echo.labstack.com/">Echo (Go) (server-side)</a>.
        </p>

        <p class="mb-6 leading-normal">
            For more information about how to use <a class="text-indigo-600 underline hover:text-orange-500" href="https://svelte.dev/">Svelte</a>
            as <a class="text-indigo-600 underline hover:text-orange-500" href="https://inertiajs.com/client-side-setup">client-side</a>
            and <a class="text-indigo-600 underline hover:text-orange-500" href="https://github.com/elipZis/inertia-echo">Echo (Go)</a>
            as <a class="text-indigo-600 underline hover:text-orange-500" href="https://inertiajs.com/server-side-setup">server-side</a>
            check the according <a class="text-indigo-600 underline hover:text-orange-500" href="https://inertiajs.com/installation">documentation</a>.
        </p>

        <p class="mb-12 leading-normal italic">
            This is an example app to illustrate the basic usage of different libraries. This demo, source and the whole package comes
            without warranty. It may or may not harm your computer.<br />
            It is not a reference for best-practices or security concerns or any other application concept but an exemplary use-case on how
            to integrate multiple systems together.<br />
            Please use with care and not as absolute reference. Any damage cannot be related back to the author.
        </p>

        <div>
            <InertiaLink class="btn-indigo mr-1" href="/500">
                500 error
            </InertiaLink>

            <InertiaLink class="btn-indigo" href="/404">
                404 error
            </InertiaLink>
        </div>
    </div>
</Layout>
