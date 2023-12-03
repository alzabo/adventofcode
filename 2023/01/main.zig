const std = @import("std");
const info = std.log.info;
const parseInt = std.fmt.parseInt;
const expect = std.testing.expect;
const assert = std.debug.assert;
const GenericReader = std.io.GenericReader;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const in_stream = buf_reader.reader();
    try part1(in_stream);

    try file.seekTo(0);

    try part2(in_stream);
}

fn part1(reader: anytype) !void {
    var buf: [1024]u8 = undefined;
    var ct1: u32 = 0;

    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var vals: [2]u8 = [_]u8{ 0, 0 };
        for (line) |c| {
            switch (c) {
                '0'...'9' => {
                    if (vals[0] == 0) {
                        vals[0] = c;
                    }
                    vals[1] = c;
                },
                else => {},
            }
        }
        //info("parsing {s}...", .{vals});
        ct1 += try parseInt(u8, &vals, 10);
    }
    info("part1: {d}", .{ct1});
}

fn part2(reader: anytype) !void {
    var buf: [1024]u8 = undefined;
    var ct1: u32 = 0;

    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var vals: [2]u8 = [_]u8{ 0, 0 };
        //info("line {s}", .{line});
        for (line, 0..) |c, i| {
            var v: u8 = 0;
            //info("char: {u}", .{c});
            switch (c) {
                '0'...'9' => {
                    v = c;
                },
                // one, two, three, four, five, six, seven, eight, nine
                // TODO: More elegant way of doing this
                'e' => {
                    if (in_line(i, line, "eight")) {
                        v = '8';
                    }
                },
                'f' => {
                    if (in_line(i, line, "four")) {
                        v = '4';
                    }
                    if (in_line(i, line, "five")) {
                        v = '5';
                    }
                },
                'n' => {
                    if (in_line(i, line, "nine")) {
                        v = '9';
                    }
                },
                'o' => {
                    if (in_line(i, line, "one")) {
                        v = '1';
                    }
                },
                's' => {
                    if (in_line(i, line, "six")) {
                        v = '6';
                    }
                    if (in_line(i, line, "seven")) {
                        v = '7';
                    }
                },
                't' => {
                    if (in_line(i, line, "two")) {
                        v = '2';
                    }
                    if (in_line(i, line, "three")) {
                        v = '3';
                    }
                },
                else => {},
            }
            if (v == 0) {
                continue;
            }
            if (vals[0] == 0) {
                vals[0] = v;
            }
            vals[1] = v;
            //info("vals 0: {u}, 1: {u}", .{ vals[0], vals[1] });
        }
        //info("parsing {s}...", .{vals});
        ct1 += try parseInt(u8, &vals, 10);
    }
    info("part2: {d}", .{ct1});
}

fn in_line(offset: usize, line: []const u8, needle: []const u8) bool {
    if (line.len - offset < needle.len) {
        return false;
    }
    return (std.mem.eql(u8, line[offset .. offset + needle.len], needle));
}

fn str_eql(a: []const u8, b: []const u8) bool {
    return (std.mem.eql(u8, a, b));
}

test "comparing slices" {
    assert(in_line(0, "seven1234foo", "seven"));
    assert(in_line(1, "1seven1234foo", "seven"));
    assert(in_line(10, "1seven1234foo", "foo"));

    //try expect(line[0..seven.len] == seven);
}
