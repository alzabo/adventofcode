const std = @import("std");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();
    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;

    var last1: u32 = 0;
    var ct1: u32 = 0;

    const ArrayList = std.ArrayList;

    var values = ArrayList(u32).init(allocator);
    defer values.deinit();

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const v = std.fmt.parseInt(u32, line, 10);
        // TODO: Review unwrapping error unions.
        // Not real confident that these are being handled as well as they could be
        if (v) |val| {
            if (last1 > 0 and val > last1) {
                ct1 += 1;
                //std.log.info("{d} gt {d}", .{ val, last });
            }
            last1 = val;
            try values.append(val);
        } else |_| {
            unreachable;
        }
        //std.log.info("line {s}", .{line});
    }
    std.log.info("part1 count: {}", .{ct1});

    var last2: u32 = 0;
    var ct2: u32 = 0;
    var i: usize = 0;
    while (true) : (i += 1) {
        //std.log.info("iteration {d}", .{i});
        if (i + 3 > values.items.len) {
            break;
        }
        var sum: u32 = 0;
        for (values.items[i .. i + 3]) |v| {
            sum += v;
        }
        //std.log.info("sum {d}", .{sum});
        if (last2 > 0 and sum > last2) {
            ct2 += 1;
        }
        last2 = sum;
    }
    std.log.info("part2 count: {}", .{ct2});
}
