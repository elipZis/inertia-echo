<script>
    import { Inertia } from '@inertiajs/inertia';
    import { InertiaLink, page } from '@inertiajs/inertia-svelte';
    import Helmet from '@/Shared/Helmet.svelte';
    import Layout from '@/Shared/Layout.svelte';
    import LoadingButton from '@/Shared/LoadingButton.svelte';
    import TextInput from '@/Shared/TextInput.svelte';
    import SelectInput from '@/Shared/SelectInput.svelte';

    const route = window.route;

    $: errors = $page.errors ?? [];

    let sending = false;
    let values = {
        Name: '',
        Email: '',
        Phone: '',
        Address: '',
        City: '',
        Region: '',
        Country: '',
        PostalCode: ''
    };

    function handleChange({ target: { name, value } }) {
        values = {
            ...values,
            [name]: value
        };
    }

    function handleSubmit() {
        sending = true;
        Inertia.post(route('organizations.store'), values).then(() => sending = false);
    }
</script>

<Helmet title="Create Organization" />

<Layout>
    <div>
        <h1 class="mb-8 font-bold text-3xl">
            <InertiaLink
                href={route('organizations')}
                class="text-indigo-600 hover:text-indigo-700"
            >
                Organizations
            </InertiaLink>

            <span class="text-indigo-600 font-medium"> /</span> Create
        </h1>

        {#if errors && errors["general"] }
            <div class="form-error">{errors["general"]}</div>
        {/if}

        <div class="bg-white rounded shadow overflow-hidden max-w-3xl">
            <form on:submit|preventDefault={handleSubmit}>
                <div class="p-8 -mr-6 -mb-8 flex flex-wrap">
                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Name*"
                        name="Name"
                        errors={errors["Organization.Name"]}
                        value={values.Name}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Email"
                        name="Email"
                        type="email"
                        errors={errors["Organization.Email"]}
                        value={values.Email}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Phone"
                        name="Phone"
                        type="text"
                        errors={errors["Organization.Phone"]}
                        value={values.Phone}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Address"
                        name="Address"
                        type="text"
                        errors={errors["Organization.Address"]}
                        value={values.Address}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="City"
                        name="City"
                        type="text"
                        errors={errors["Organization.City"]}
                        value={values.City}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Province/State"
                        name="Region"
                        type="text"
                        errors={errors["Organization.Region"]}
                        value={values.Region}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Country"
                        name="Country"
                        errors={errors["Organization.Country"]}
                        value={values.Country}
                        onChange={handleChange}
                    >
                        <option value=""></option>
                        <option value="CA">Canada</option>
                        <option value="US">United States</option>
                    </SelectInput>

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Postal Code"
                        name="PostalCode"
                        type="text"
                        errors={errors["Organization.PostalCode"]}
                        value={values.PostalCode}
                        onChange={handleChange}
                    />
                </div>

                <div class="px-8 py-4 bg-gray-100 border-t border-gray-200 flex justify-end items-center">
                    <LoadingButton
                        loading={sending}
                        type="submit"
                        className="btn-indigo"
                    >
                        Create Organization
                    </LoadingButton>
                </div>
            </form>
        </div>
    </div>
</Layout>
