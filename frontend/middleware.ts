import { NextResponse, NextRequest } from "next/server";

export function middleware(request: NextRequest) {
    const headers: Headers = new Headers(request.headers);
    headers.set("x-current-path", request.nextUrl.pathname);

    return NextResponse.next({ headers });
}
