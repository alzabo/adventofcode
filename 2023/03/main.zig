const std = @import("std");
const info = std.log.info;
const parseInt = std.fmt.parseInt;
const expect = std.testing.expect;
const assert = std.debug.assert;
const ArenaAllocator = std.heap.ArenaAllocator;

// TODO: What is this number and where does it come from?
const magic_invalid_number: u8 = 0b10101010;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    var matrix: [140][140]u8 = undefined;

    var i: usize = 0;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        std.mem.copy(u8, &matrix[i], line);
        i += 1;
    }

    var arena = ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    var map = std.AutoHashMap([2]usize, u16).init(allocator);

    for (matrix, 0..) |row, x| {
        for (row, 0..) |c, y| {
            switch (c) {
                '0'...'9', '.', '\n', magic_invalid_number => {},
                else => {
                    //std.debug.print("x: {d} y: {d} m: {c} c: {c}\n", .{ x, y, matrix[x][y], c });
                    // TODO: Get a better sense of how to do this properly rather than just
                    // messing with it until it works
                    const _x: u8 = @truncate(x);
                    assert(x == _x);
                    const _y: u8 = @truncate(y);
                    assert(y == _y);
                    //std.debug.print("x: {d} y: {d} _x: {d} _y: {d} m: {c} c: {c}\n", .{ x, y, _x, _y, matrix[x][y], c });

                    try find_neighboring_numbers(_x, _y, matrix, &map);
                },
            }
        }
    }

    var ct1: u32 = 0;
    var iter = map.iterator();
    while (iter.next()) |kv| {
        ct1 += kv.value_ptr.*;
    }
    info("part1: {d}", .{ct1});

    var ct2: u64 = 0;
    for (matrix, 0..) |row, x| {
        for (row, 0..) |c, y| {
            switch (c) {
                '*' => {
                    //std.debug.print("x: {d} y: {d} m: {c} c: {c}\n", .{ x, y, matrix[x][y], c });
                    // TODO: Get a better sense of how to do this properly rather than just
                    // messing with it until it works
                    const _x: u8 = @truncate(x);
                    assert(x == _x);
                    const _y: u8 = @truncate(y);
                    assert(y == _y);
                    //std.debug.print("x: {d} y: {d} _x: {d} _y: {d} m: {c} c: {c}\n", .{ x, y, _x, _y, matrix[x][y], c });

                    ct2 += try product_of_two_neighboring_numbers(_x, _y, matrix);
                },
                else => {},
            }
        }
    }
    info("part2: {d}", .{ct2});
}

fn product_of_two_neighboring_numbers(x: i16, y: i16, matrix: anytype) !u32 {
    var seen: [2][2]usize = undefined;
    var numbers: [2]u32 = [_]u32{ 0, 0 };
    var nindex: usize = 0;

    coords: for (neighbor_coords) |d| {
        if (x + d[0] < 0 or y + d[1] < 0) {
            continue;
        }
        const ux: usize = @intCast(x + d[0]);
        const uy: usize = @intCast(y + d[1]);
        if (ux > matrix.len or uy > matrix[0].len) {
            continue;
        }

        var buf: [3]u8 = [_]u8{ 0, 0, 0 };
        var bufidx: usize = 0;
        //std.debug.print("considering... ux: {d} uy: {d}; \n", .{ ux, uy });
        switch (matrix[ux][uy]) {
            '0'...'9' => {
                const _x: usize = ux;
                var _y: usize = uy;
                //std.debug.print("_x: {d}; _y: {d}\n", .{ _x, _y });
                // find the first digit
                while (_y > 0) {
                    switch (matrix[_x][_y - 1]) {
                        '0'...'9' => {
                            _y -= 1;
                        },
                        else => {
                            break;
                        },
                    }
                }
                const start = [_]usize{ _x, _y };
                for (seen) |v| {
                    if (v[0] == start[0] and v[1] == start[1]) {
                        continue :coords;
                    }
                }
                //std.debug.print("_x: {d}; _y: {d}\n", .{ _x, _y });

                while (_y < matrix[0].len) : (_y += 1) {
                    const val = matrix[_x][_y];
                    switch (val) {
                        '0'...'9' => {
                            buf[bufidx] = val;
                            bufidx += 1;
                        },
                        else => {
                            break;
                        },
                    }
                }
                const end: usize = std.mem.indexOfScalar(u8, &buf, 0) orelse buf.len;

                //std.debug.print("buf: {b}\n", .{buf[0..end]});
                const value = try parseInt(u32, buf[0..end], 10);
                buf = [_]u8{ 0, 0, 0 };
                bufidx = 0;

                //std.debug.print("value: {d} from: x: {d} y: {d}\n", .{ value, x, y });
                numbers[nindex] = value;
                seen[nindex] = start;
                nindex += 1;
            },
            else => {},
        }
    }
    //std.debug.print("values: {d} from: x: {d} y: {d}\n", .{ numbers, x, y });
    if (numbers[0] == 0 or numbers[1] == 0) {
        return 0;
    }
    return numbers[0] * numbers[1];
}

fn find_neighboring_numbers(x: i16, y: i16, matrix: anytype, map: anytype) !void {
    for (neighbor_coords) |d| {
        const ix = x + d[0];
        const iy = y + d[1];
        if (ix < 0 or iy < 0) {
            continue;
        }
        const ux: usize = @intCast(ix);
        const uy: usize = @intCast(iy);
        if (ux > matrix.len or uy > matrix[0].len) {
            continue;
        }

        var buf: [3]u8 = [_]u8{ 0, 0, 0 };
        var bufidx: usize = 0;
        //std.debug.print("considering... ux: {d} uy: {d}; \n", .{ ux, uy });
        switch (matrix[ux][uy]) {
            '0'...'9' => {
                const _x: usize = ux;
                var _y: usize = uy;
                //std.debug.print("_x: {d}; _y: {d}\n", .{ _x, _y });
                // find the first digit
                while (_y > 0) {
                    switch (matrix[_x][_y - 1]) {
                        '0'...'9' => {
                            _y -= 1;
                        },
                        else => {
                            break;
                        },
                    }
                }
                const start = [_]usize{ _x, _y };
                if (map.get(start)) |_| {
                    continue;
                }
                //std.debug.print("_x: {d}; _y: {d}\n", .{ _x, _y });

                while (_y < matrix[0].len) : (_y += 1) {
                    const val = matrix[_x][_y];
                    switch (val) {
                        '0'...'9' => {
                            buf[bufidx] = val;
                            bufidx += 1;
                        },
                        else => {
                            break;
                        },
                    }
                }
                const end: usize = std.mem.indexOfScalar(u8, &buf, 0) orelse buf.len;

                //std.debug.print("buf: {b}\n", .{buf[0..end]});
                const value = try parseInt(u16, buf[0..end], 10);
                buf = [_]u8{ 0, 0, 0 };
                bufidx = 0;

                std.debug.print("value: {d} x: {d} y: {d} from: x: {d} y: {d}\n", .{ value, start[0], start[1], x, y });
                try map.put(start, value);
            },
            else => {},
        }
    }
}

const neighbor_coords = [_][2]i32{
    [_]i32{ 1, 0 }, // down
    [_]i32{ -1, 0 }, // up
    [_]i32{ 0, 1 }, // right
    [_]i32{ 0, -1 }, // left
    [_]i32{ 1, 1 }, // down, right
    [_]i32{ 1, -1 }, // down, left
    [_]i32{ -1, 1 }, // up, right
    [_]i32{ -1, -1 }, // up, left
};

test "lookaround" {
    const m = [3][8]u8{
        [_]u8{ '1', '2', '3', '.', '.', '.', '8', '.' },
        [_]u8{ '.', '%', '.', '.', '=', '.', '.', '7' },
        [_]u8{ '.', '.', '.', '4', '5', '6', '.', '.' },
    };
    var map = std.AutoHashMap([2]usize, u16).init(std.testing.allocator);
    defer map.deinit();

    //var iter = map.iterator();
    //while (iter.next()) |kv| {
    //    std.debug.print("x: {d} y: {d} val: {d}\n", .{ kv.key_ptr[0], kv.key_ptr[1], kv.value_ptr });
    //}

    try find_neighboring_numbers(1, 1, m, &map);
    if (map.get([2]usize{ 0, 0 })) |r| {
        try expect(r == 123);
    } else {
        unreachable;
    }

    try find_neighboring_numbers(1, 4, m, &map);
    if (map.get([2]usize{ 2, 3 })) |r| {
        try expect(r == 456);
    } else {
        unreachable;
    }

    if (map.get([2]usize{ 0, 8 })) |_| {
        unreachable;
    }

    var iter = map.iterator();
    while (iter.next()) |kv| {
        std.debug.print("x: {d} y: {d} val: {d}\n", .{ kv.key_ptr[0], kv.key_ptr[1], kv.value_ptr });
    }
}
