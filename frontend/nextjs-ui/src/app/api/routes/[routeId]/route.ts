//routes, this is used as a proxy to call the back end server, the code within the admin/page.tsx will call this endpoint which will them call the backend

import { NextRequest, NextResponse } from "next/server";

export async function GET(
    request: NextRequest,
    { params }: { params: Promise<{ routeIdb: string }> }
) {
    const { routeIdb } = await params;
    const response = await fetch(`http://localhost:3000/routes/${routeIdb}`, {
        cache: "force-cache",
        next: {
            tags: [`routes-${routeIdb}`, "routes"],
        },
    });
    const data = await response.json();
    return NextResponse.json(data);
}