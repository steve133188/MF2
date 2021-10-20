import * as React from 'react';
import Typography from '@mui/material/Typography';
import Pagination from '@mui/material/Pagination';
import Stack from '@mui/material/Stack';

export function PaginationControlled() {
    const [page, setPage] = React.useState(1);
    const handleChange = (event, value) => {
        setPage(value);
    };

    return (
        <div className="pagination">
            <Stack spacing={2}>
                <Pagination count={10} page={page} onChange={handleChange} shape="rounded"/>
            </Stack>
        </div>
    );
}
