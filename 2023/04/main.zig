const std = @import("std");
const info = std.log.info;
const parseInt = std.fmt.parseInt;
const expect = std.testing.expect;
const assert = std.debug.assert;
const ArenaAllocator = std.heap.ArenaAllocator;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    var ct1: u32 = 0;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        // Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        var winner = std.mem.zeroes([10]u8);
        var wIndex: u8 = 0;

        var numbers = std.mem.zeroes([25]u8);
        var nIndex: u8 = 0;

        var inWin: bool = false;
        var inNum: bool = false;
        var inVal: bool = false;
        var valBuf: [3]u8 = [_]u8{ 0, 0, 0 };
        var valIdx: u8 = 0;

        for (line) |c| {
            //info("c: '{c}' inWin: {any} inVal {any}", .{ c, inWin, inVal });
            switch (c) {
                ':' => {
                    inWin = true;
                },
                '|' => {
                    inWin = false;
                    inNum = true;
                },
                ' ' => {
                    if (inVal) {
                        //info("valBuf: '{s}'", .{valBuf});
                        const val = try parseIntBuf(&valBuf);
                        valBuf = [_]u8{ 0, 0, 0 };
                        valIdx = 0;
                        //info("val: {d}", .{val});

                        if (inWin) {
                            winner[wIndex] = val;
                            wIndex += 1;
                        }

                        if (inNum) {
                            numbers[nIndex] = val;
                            nIndex += 1;
                        }
                    }
                    inVal = false;
                },
                '0'...'9' => {
                    if (!(inWin or inNum)) {
                        continue;
                    }
                    inVal = true;
                    valBuf[valIdx] = c;
                    valIdx += 1;
                },
                else => {},
            }

            //info("c: '{c}'", .{c});
        }
        // parse the final number not handled in the switch statement
        numbers[nIndex] = try parseIntBuf(&valBuf);

        var score: u32 = 0;
        const e = std.mem.indexOfScalar(u8, &winner, 0) orelse winner.len;
        for (winner[0..e]) |w| {
            for (numbers) |n| {
                if (w == n) {
                    if (score == 0) {
                        score += 1;
                    } else {
                        score *= 2;
                    }
                }
            }
        }
        //info("matches: {d} val: {d}", .{ matches, val });

        ct1 += score;
        //info("vals: {d}", .{winner});
    }
    info("part1: {d}", .{ct1});
}

fn parseIntBuf(b: []const u8) !u8 {
    const end = std.mem.indexOfScalar(u8, b, 0) orelse b.len;
    const val = try parseInt(u8, b[0..end], 10);
    return val;
}

const Card = struct {
    win: [10]u8,
    nums: [25]u8,
};
